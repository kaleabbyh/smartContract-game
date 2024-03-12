package models

import (
	"errors"
	"time"

	uuid "github.com/google/uuid"
	"gorm.io/gorm"
)

type Model struct {
	ID        string     `gorm:"primary_key" json:"id,omitempty"`
	CreatedAt time.Time  `gorm:"not null" json:"created_at" sql:"DEFAULT:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time  `gorm:"not null" json:"updated_at" sql:"DEFAULT:CURRENT_TIMESTAMP"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at,omitempty"`
}

func (model *Model) BeforeCreate(scope *gorm.DB) error {
	id, err := uuid.New().MarshalText()
	if err == nil {
		model.ID = string(id)
		return nil
	}

	return errors.New("unable to create id")
}
