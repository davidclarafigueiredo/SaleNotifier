package appactions

import (
	"encoding/json"
	"os"

	"github.com/davidclarafigueiredo/SaleNotifier/actions"
	"github.com/davidclarafigueiredo/SaleNotifier/connect"
	"github.com/davidclarafigueiredo/SaleNotifier/handler"
	"github.com/davidclarafigueiredo/SaleNotifier/scraper"
	"github.com/rs/zerolog/log"
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
	if actions.ComparePrice(url, apiUrl) {
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
	if actions.ComparePrice(url, apiUrl) {
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
