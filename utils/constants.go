package utils

import "gorm.io/gorm"

type ServerConfigs struct {
	DB              *gorm.DB `json:"db"`
	SecretKey       string   `json:"secret_key"`
	NEXTAUTH_SECRET string   `json:"nextauth"`
}
