package appactions

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/davidclarafigueiredo/SaleNotifier/connect"
	"github.com/davidclarafigueiredo/SaleNotifier/handler"
	"github.com/davidclarafigueiredo/SaleNotifier/scraper"
	"github.com/rs/zerolog/log"
	"github.com/shopspring/decimal"
)

type GameStruct struct {
	Nsuid           string `json:"Nsuid"`
	Url             string `json:"Url"`
	ApiUrl          string `json:"ApiUrl"`
	GameTitle       string `json:"GameTitle"`
	IsDiscounted    string `json:"IsDiscounted"`
	RegularPrice    string `json:"RegularPrice"`
	DiscountedPrice string `json:"DiscountedPrice"`
}

//export GetInformation
func GetInformation(url string) string {
	nsuid := scraper.GetNSUID(url)
	apiUrl := "https://api.ec.nintendo.com/v1/price?country=DE&lang=de&ids=" + nsuid

	gameTitle := scraper.GetGameTitle(url)
	regularPrice := scraper.GetPrice(url)
	discountedPrice := handler.GetPrice(connect.Connect(apiUrl))

	isDiscounted := "not on sale"
	if comparePrice(url, apiUrl) {
		isDiscounted = "on sale"
	}

	return gameTitle + " ;" + regularPrice + " ;" + discountedPrice + " ;" + isDiscounted

}

//export WriteToJSON
func WriteEntryToJSON(jsonFileName string, url string) {

	// Open the json file
	err := checkFile(jsonFileName)
	if err != nil {
		log.Fatal().Err(err).Msg("Error opening/creating output file")
	}
	jsonFile, err := os.ReadFile(jsonFileName)
	if err != nil {
		log.Fatal().Err(err).Msg("Error opening output file for reading")
	}

	// Read the json file
	data := []GameStruct{}
	json.Unmarshal(jsonFile, &data)

	// Create a new game struct
	nsuid := scraper.GetNSUID(url)
	apiUrl := "https://api.ec.nintendo.com/v1/price?country=DE&lang=de&ids=" + nsuid
	gameTitle := scraper.GetGameTitle(url)
	regularPrice := scraper.GetPrice(url)
	discountedPrice := handler.GetFormPrice(connect.Connect(apiUrl))
	isDiscounted := "not on sale"
	if comparePrice(url, apiUrl) {
		isDiscounted = "on sale"
	}

	newGame := &GameStruct{
		Nsuid:           nsuid,
		Url:             url,
		ApiUrl:          apiUrl,
		GameTitle:       gameTitle,
		IsDiscounted:    isDiscounted,
		RegularPrice:    regularPrice,
		DiscountedPrice: discountedPrice,
	}

	// Append the new game to the data if the title is not already in the list
	for _, game := range data {
		if game.GameTitle == newGame.GameTitle {
			return
		}
	}
	data = append(data, *newGame)

	// Preparing the data to be marshalled and written.
	dataBytes, err := json.Marshal(data)
	if err != nil {
		log.Fatal().Err(err).Msg("Error marshalling data")
	}
	// Write data to the json file
	err = os.WriteFile(jsonFileName, dataBytes, 0644)
	if err != nil {
		log.Fatal().Err(err).Msg("Error writing to output file")
	}
}

//export UpdateJSONEntry
func UpdateJSONEntry(jsonFileName string, url string) bool {
	// Open the json file
	err := checkFile(jsonFileName)
	if err != nil {
		log.Fatal().Err(err).Msg("Error opening/creating output file")
	}
	jsonFile, err := os.ReadFile(jsonFileName)
	if err != nil {
		log.Fatal().Err(err).Msg("Error opening output file for reading")
	}

	// Get sale status from http request
	nsuid := scraper.GetNSUID(url)
	apiUrl := "https://api.ec.nintendo.com/v1/price?country=DE&lang=de&ids=" + nsuid
	isDiscountedNew := "not on sale"
	if comparePrice(url, apiUrl) {
		isDiscountedNew = "on sale"
	}

	// Read the json file
	data := []GameStruct{}
	json.Unmarshal(jsonFile, &data)

	var isDiscountedOld string

	// Get sale status in json file und update it
	for i, game := range data {
		if game.Nsuid == nsuid {
			isDiscountedOld = game.IsDiscounted
			if game.IsDiscounted != isDiscountedNew {
				data[i].IsDiscounted = isDiscountedNew
			}
		}
	}

	// Preparing the data to be marshalled and written.
	dataBytes, err := json.Marshal(data)
	if err != nil {
		log.Fatal().Err(err).Msg("Error marshalling data")
	}
	// Write data to the json file
	err = os.WriteFile(jsonFileName, dataBytes, 0644)
	if err != nil {
		log.Fatal().Err(err).Msg("Error writing to output file")
	}

	if isDiscountedOld == "not on sale" && isDiscountedNew == "on sale" {
		return true // discount states changed to on sale
	}

	return false
}

//export RemoveEntryFromJSON
func RemoveEntryFromJSON(jsonFileName string, nsuid string) {
	// Open the json file
	err := checkFile(jsonFileName)
	if err != nil {
		log.Fatal().Err(err).Msg("Error opening/creating output file")
	}
	jsonFile, err := os.ReadFile(jsonFileName)
	if err != nil {
		log.Fatal().Err(err).Msg("Error opening output file for reading")
	}

	// Read the json file
	data := []GameStruct{}
	json.Unmarshal(jsonFile, &data)

	// Remove the game from the data
	for i, game := range data {
		if game.Nsuid == nsuid {
			data = append(data[:i], data[i+1:]...)
			break
		}
	}

	// Preparing the data to be marshalled and written.
	dataBytes, err := json.Marshal(data)
	if err != nil {
		log.Fatal().Err(err).Msg("Error marshalling data")
	}
	// Write data to the json file
	err = os.WriteFile(jsonFileName, dataBytes, 0644)
	if err != nil {
		log.Fatal().Err(err).Msg("Error writing to output file")
	}
}

func checkFile(filename string) error {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		_, err := os.Create(filename)
		if err != nil {
			return err
		}
	}
	return nil
}

// returns a boolean if game is part of the wishlist (i think?)
// func contains(list []string, item string) bool {
// 	for _, v := range list {
// 		if v == item {
// 			return true
// 		}
// 	}
// 	return false
// }

func comparePrice(url string, apiUrl string) bool {

	price, _ := decimal.NewFromString(handler.GetPrice(connect.Connect(apiUrl)))
	discountPrice, _ := decimal.NewFromString(scraper.GetPrice(url))

	if !price.Equal(discountPrice) {
		fmt.Println("Price: ", price)
		fmt.Println("Discount Price: ", discountPrice)
		// fmt.Printf("%s is on sale for %s", scraper.GetGameTitle(url), handler.GetFormPrice(connect.Connect(apiUrl)))
		return true
	}
	return false
}
