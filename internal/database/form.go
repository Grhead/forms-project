package database

import (
	"time"
	"tusur-forms/internal/domain"

	"gorm.io/gorm"
)

type dbForm struct {
	Id            string `gorm:"primaryKey"`
	Title         string
	DocumentTitle string
	ExternalId    string
	CreatedAt     time.Time
}

func CreateForm(f *domain.Form, db *gorm.DB) error {
	dbF := dbForm{
		Id:         f.Id,
		ExternalId: f.ExternalId,
		CreatedAt:  f.CreatedAt,
	}

	err := db.Create(&dbF).Error
	if err != nil {
		return err
	}
	return db.Save(&dbF).Error
}
