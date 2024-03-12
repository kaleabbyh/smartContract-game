package repositories

import (
	"cicada/web-service-gin/models"
	"errors"

	"gorm.io/gorm"
)

type GameRequest struct {
	PlayerScore []models.PlayerScore
	Players     []models.Poolli
}

func CreateNePlayer(db *gorm.DB, playerScore models.PlayerScore) (models.PlayerScore, error) {
	if err := db.Create(&playerScore).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return models.PlayerScore{}, errors.New("duplicated key")
		}
		return models.PlayerScore{}, err
	}
	return playerScore, nil
}

func GetPlayerScoreByID(db *gorm.DB, id string) (*models.PlayerScore, error) {

	var PlayerScore models.PlayerScore
	if err := db.First(&PlayerScore, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &PlayerScore, nil
}

func UpdatePlayerScore(db *gorm.DB, id string, updatedPlayerScore *models.PlayerScore) (*models.PlayerScore, error) {

	if err := db.Save(&updatedPlayerScore).Error; err != nil {
		return nil, err
	}

	return updatedPlayerScore, nil
}

func GetPlayerScoreByGameId(db *gorm.DB, gameid string) ([]models.PlayerScore, error) {
	var playerScore []models.PlayerScore

	if err := db.Find(&playerScore, "gameid = ?", gameid).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("player score not found")
		}

		return nil, err
	}
	return playerScore, nil
}

func GetWinnerForGame(db *gorm.DB, gameid string) (GameRequest, error) {
	var playerScore []models.PlayerScore

	if err := db.Find(&playerScore, "gameid = ?", gameid).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return GameRequest{}, errors.New("player score not found")
		}

		return GameRequest{}, err
	}

	var pool models.Pool
	if err := db.First(&pool, "gameid = ?", gameid).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return GameRequest{}, errors.New("pool not found")
		}

		return GameRequest{}, err
	}

	var poollis []models.Poolli

	if err := db.Find(&poollis, "pool_id = ?", pool.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return GameRequest{}, errors.New("pool-li not found")
		}

		return GameRequest{}, err
	}

	gameData := GameRequest{
		PlayerScore: playerScore,
		Players:     poollis,
	}
	return gameData, nil
}

func UpdatePlayerScoreGameId(db *gorm.DB, oldGameID string, newGameID string) ([]models.PlayerScore, error) {

	var playerScores []models.PlayerScore
	newUpdate := map[string]interface{}{"gameid": newGameID}

	result := db.Model(&models.PlayerScore{}).Where("gameid = ?", oldGameID).Updates(newUpdate)
	if result.Error != nil {
		return nil, result.Error
	}

	if err := db.Find(&playerScores, "gameid = ?", newGameID).Error; err != nil {
		return nil, err
	}

	return playerScores, nil
}
