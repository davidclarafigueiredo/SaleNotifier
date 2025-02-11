package scraper

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
)

func GetURL() string {
	file, err := os.Open("data/import")
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	if scanner.Scan() { // Reads the first line
		return scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
	return ""
}

func GetInfo(info string) string {
	url := GetURL()
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

func GetNSUID() string {
	return GetInfo("offdeviceNsuID")
}

func GetGameTitle() string {
	return GetInfo("gameTitle")
}
