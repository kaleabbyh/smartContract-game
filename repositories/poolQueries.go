package repositories

import (
	"cicada/web-service-gin/models"
	"errors"

	"gorm.io/gorm"
)

func CreateNewPool(db *gorm.DB, newPool *models.Pool) (*models.Pool, error) {
	if err := db.Create(&newPool).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, errors.New("duplicated key")
		}
		return nil, err
	}
	return newPool, nil
}

func UpdatePool(db *gorm.DB, id string, updatedPool models.Pool) (models.Pool, error) {
	var existingPool models.Pool

	if err := db.First(&existingPool, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Pool{}, errors.New("record not found")
		}
		return models.Pool{}, err
	}

	if updatedPool.Gameid != "" {
		existingPool.Gameid = updatedPool.Gameid
	}

	if updatedPool.GameStatus != "" {
		existingPool.GameStatus = updatedPool.GameStatus
	}

	if err := db.Save(&existingPool).Error; err != nil {
		return models.Pool{}, err
	}

	return existingPool, nil
}

func CreateNewPoolli(db *gorm.DB, newPooli *models.Poolli) (*models.Poolli, error) {
	if err := db.Create(&newPooli).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, errors.New("duplicated key")
		}
		return nil, err
	}
	return newPooli, nil
}

func GetAllPools(db *gorm.DB) ([]models.Pool, error) {
	var pools []models.Pool

	if err := db.Preload("Poolli").Find(&pools).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("no pools found")
		}

		return nil, err
	}
	return pools, nil
}

func GetPoolByID(db *gorm.DB, id string) (models.Pool, error) {
	var pool models.Pool

	if err := db.Preload("Poolli").First(&pool, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Pool{}, errors.New("no pool found")
		}

		return models.Pool{}, err
	}
	return pool, nil
}

func DeletePoolWithPoollisByID(db *gorm.DB, poolID string) error {
	tx := db.Begin()

	if err := tx.Where("pool_id = ?", poolID).Delete(&models.Poolli{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Where("id = ?", poolID).Delete(&models.Pool{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func GetPoollisplayerIsIn(db *gorm.DB, poolID string, player string) ([]models.Poolli, error) {
	var poollis []models.Poolli

	if err := db.Where("pool_id = ? AND ? = ANY(Owners)", poolID, player).Find(&poollis).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("no pools found")
		}

		return nil, err
	}
	return poollis, nil
}
