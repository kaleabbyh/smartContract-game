package config

import (
	"cicada/web-service-gin/models"
	"os"

	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	LoadConfig()

	var logger = log.Logger.With().Str("func", "Configs").Logger()
	// ks, err := RestoreKs()
	// if err != nil {
	// 	logger.Error().Err(err).Msgf("Error on Restoring ks: %s", err.Error())
	// }
	// _ = ks

	var DB_URL string
	DB_URL_jsonBytes, err := os.ReadFile(os.Getenv("DB_URL"))
	if err != nil {
		logger.Info().Msgf("Failed to read DB_URL file: no such file")
		DB_URL = os.Getenv("DB_URL")
	} else {
		logger.Info().Msgf("Using DB_URL_jsonBytes")
		DB_URL = string(DB_URL_jsonBytes)
	}
	db, err := gorm.Open(postgres.Open(DB_URL), &gorm.Config{})
	if err != nil {
		logger.Error().Err(err).Msgf("Error in DB connection: %s", err.Error())
		return nil
	}

	logger.Info().Msgf("DB connected successfully")
	return db
}

func Migrate() {
	logger := log.Logger.With().Str("func", "MigrateTables").Logger()
	LoadConfig()
	DB := ConnectDB()
	err := DB.AutoMigrate(
		models.PlayerScore{},
		models.Pool{},
		models.Poolli{},
		models.AdminWallet{},
		models.EscrowHoF{},
	)
	if err != nil {
		logger.Error().Err(err).Msgf("Error migrating tables: %s", err.Error())
		return
	}

}
