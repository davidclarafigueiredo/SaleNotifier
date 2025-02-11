package scraper

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
)

func GetInfo(info string, url string) string {
	if url == "" {
		log.Fatal("No URL found in wishlist.txt")
	}

	// Send HTTP request
	response, err := http.Get(url)
	if err != nil {
		log.Fatal("Error fetching the webpage:", err)
	}
	defer response.Body.Close()

	// Read response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal("Error reading response body:", err)
	}

	titleRegex := regexp.MustCompile(`"` + info + `"\s*:\s*"([^"]+)"`)
	matches := titleRegex.FindStringSubmatch(string(body))

	if len(matches) > 1 {
		return matches[1]
	}

	fmt.Printf("%s not found.\n", info)
	return ""
}

func GetNSUID(url string) string {
	return GetInfo("offdeviceNsuID", url)
}

func GetGameTitle(url string) string {
	return GetInfo("gameTitle", url)
}

func GetPrice(url string) string {
	return GetInfo("offdeviceProductPrice", url)
}
