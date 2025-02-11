package scraper

import (
	"bufio"
	"io"
	"net/http"
	"os"
	"regexp"

	"github.com/rs/zerolog/log"
)

func GetURL() string {
	file, err := os.Open("data/import")
	if err != nil {
		log.Error().Err(err).Msg("Error opening file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	if scanner.Scan() { // Reads the first line
		return scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Error().Err(err).Msg("Error reading file")
	}
	log.Fatal().Msg("URL not found")
	return ""
}

func GetInfo(info string) string {
	url := GetURL()

	response, err := http.Get(url)
	if err != nil {
		log.Fatal().Err(err).Msg("Error fetching the webpage")
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal().Err(err).Msg("Error reading response body")
	}

	titleRegex := regexp.MustCompile(`"` + info + `"\s*:\s*"([^"]+)"`)
	matches := titleRegex.FindStringSubmatch(string(body))

	if len(matches) > 1 {
		return matches[1]
	}

	log.Info().Msgf("%s not found.", info)
	return ""
}

func GetNSUID() string {
	return GetInfo("offdeviceNsuID")
}

func GetGameTitle() string {
	return GetInfo("gameTitle")
}
