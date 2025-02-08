package config

import (
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func Init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Error().Err(err).Msg("Could not load .env file")
	}
}
