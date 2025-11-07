package database

import (
	"time"
	"tusur-forms/internal/domain"

	"gorm.io/gorm"
)

type dbForm struct {
	ID            string `gorm:"primaryKey"`
	Title         string
	DocumentTitle string
	ExternalID    string
	CreatedAt     time.Time
}

func CreateForm(f *domain.Form, db *gorm.DB) error {
	dbF := dbForm{
		ID:         f.ID,
		ExternalID: f.ExternalID,
		CreatedAt:  f.CreatedAt,
	}

	err := db.Create(&dbF).Error
	if err != nil {
		return err
	}
	return db.Save(&dbF).Error
}
