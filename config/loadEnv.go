package config

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

func LoadConfig() {
	var logger = log.Logger.With().Str("func", "Configs").Logger()
	err := godotenv.Load()
	if err != nil {
		logger.Error().Err(err).Msg("Error loading .env file:")
	}
}
