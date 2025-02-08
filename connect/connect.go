package connect

import (
	"io"
	"net/http"
	"os"

	"github.com/rs/zerolog/log"
)

func Connect() []byte {
	// Get the URL from the environment variable
	url := os.Getenv("REQUEST")
	// Connect to the site
	res, err := http.Get(url)
	if err != nil {
		log.Error().Err(err).Msg("Could not connect to site")
	}
	defer res.Body.Close()
	// Read the body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Error().Err(err).Msg("Could not read body")
	}
	return body
}
