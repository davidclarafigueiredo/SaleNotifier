package main

import (
	"github.com/davidclarafigueiredo/SaleNotifier/appactions"
	"github.com/davidclarafigueiredo/SaleNotifier/config"
	"github.com/rs/zerolog/log"
)

func main() {
	config.Init()
	// fmt.Printf("%s\n", handler.GetPrice(connect.Connect()))
	// handler.SendMail()

	// actions.SaleChecker()
	// fmt.Println(appactions.GetInformation("https://www.nintendo.com/de-de/Spiele/Nintendo-Switch-Spiele/Donkey-Kong-Country-Returns-HD-2590475.html"))

	// Testing
	filePath := "data/game_list.json"
	testUrls := [...]string{"https://www.nintendo.com/de-de/Spiele/Nintendo-Switch-Download-Software/Disney-Dreamlight-Valley-2232608.html",
		"https://www.nintendo.com/de-de/Spiele/Nintendo-Switch-Download-Software/No-Man-s-Sky-2169216.html",
		"https://www.nintendo.com/de-de/Spiele/Nintendo-Switch-Spiele/Donkey-Kong-Country-Returns-HD-2590475.html",
		"https://www.nintendo.com/de-de/Spiele/Nintendo-Switch-Spiele/Sid-Meier-s-Civilization-VII-2637632.html"}

	// for _, url := range testUrls {
	// 	appactions.WriteEntryToJSON(filePath, url)
	// }

	log.Info().Msgf("%s: %t\n", "Disney Dreamlight Valley", appactions.UpdateJSONEntry(filePath, testUrls[0]))
	log.Info().Msgf("%s: %t\n", "No Man's Sky", appactions.UpdateJSONEntry(filePath, testUrls[1]))
	log.Info().Msgf("%s: %t\n", "Donkey Kong Country Returns HD", appactions.UpdateJSONEntry(filePath, testUrls[2]))
	log.Info().Msgf("%s: %t\n", "Sid Meier's Civilization VII", appactions.UpdateJSONEntry(filePath, testUrls[3]))
}
