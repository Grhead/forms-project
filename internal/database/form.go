package database

import (
	"time"
	"tusur-forms/internal/domain"
)

type dbForm struct {
	ID            string `gorm:"primaryKey"`
	Title         string
	DocumentTitle string
	ExternalID    string
	CreatedAt     time.Time
}

func (g *GormRepository) CreateForm(f *domain.Form) error {
	dbF := dbForm{
		ID:            f.ID,
		ExternalID:    f.ExternalID,
		CreatedAt:     f.CreatedAt,
		Title:         f.Title,
		DocumentTitle: f.DocumentTitle,
	}

	err := g.db.Create(&dbF).Error
	if err != nil {
		return err
	}
	return g.db.Save(&dbF).Error
}
