package connect

import (
	"io"
	"net/http"

	"github.com/rs/zerolog/log"
)

func Connect(apiUrl string) []byte {
	res, err := http.Get(apiUrl)
	if err != nil {
		log.Error().Err(err).Msg("Could not connect to site")
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Error().Err(err).Msg("Could not read body")
	}
	log.Info().Msg("Connected to site")
	log.Info().Msg("apiUrl: " + apiUrl)
	return body
}
