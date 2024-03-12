package main

import (
	"cicada/web-service-gin/config"
	"cicada/web-service-gin/models"

	"github.com/rs/zerolog/log"
)

func main() {
	logger := log.Logger.With().Str("func", "MigrateTables").Logger()

	config.LoadConfig()
	DB := config.ConnectDB()

	err := DB.AutoMigrate(
		models.PlayerScore{},
		models.Pool{},
		models.Poolli{},
		models.AdminWallet{},
		// models.EscrowHoF{},
	)
	if err != nil {
		logger.Error().Err(err).Msgf("Error migrating tables: %s", err.Error())
		return
	}

}
