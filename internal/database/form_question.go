package database

import (
	"tusur-forms/internal/domain"

	"github.com/google/uuid"
)

type dbFormsQuestion struct {
	ID         string `gorm:"primaryKey"`
	FormID     string
	Form       dbForm `gorm:"foreignKey:FormID;references:ID"`
	QuestionID string
	Question   dbQuestion `gorm:"foreignKey:QuestionID;references:ID"`
}

func (g *GormRepository) CreateFormsQuestion(f *domain.Form, q *domain.Question) error {
	dbQ := dbFormsQuestion{
		ID:         uuid.NewString(),
		FormID:     f.ID,
		Form:       dbForm{},
		QuestionID: q.ID,
		Question:   dbQuestion{},
	}

	err := g.db.Create(&dbQ).Error
	if err != nil {
		return err
	}
	return g.db.Save(&dbQ).Error
}
