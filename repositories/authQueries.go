package repositories

import (
	"cicada/web-service-gin/models"
	"errors"
	"regexp"
	"strings"

	"gorm.io/gorm"
)

func GetAdminByWallet(db *gorm.DB, wallet string) (*models.AdminWallet, error) {
	// Check for address validation before database call
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	if !re.MatchString(wallet) {
		return nil, errors.New("invalid ethereum adddress")
	}

	var wallets models.AdminWallet
	if err := db.First(&wallets, "wallet = ?", strings.ToLower(wallet)).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("admin wallet not found")
		}
		return nil, err
	}
	return &wallets, nil
}

func CreateAdmin(db *gorm.DB, newAdmin *models.AdminWallet) (*models.AdminWallet, error) {

	if err := db.Create(&newAdmin).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, errors.New("duplicated key")
		}
		return nil, err
	}
	return newAdmin, nil
}
