package database

import (
	"tusur-forms/internal/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type dbFormsQuestion struct {
	ID         string `gorm:"primaryKey"`
	FormID     string
	Form       dbForm `gorm:"foreignKey:FormID;references:ID"`
	QuestionID string
	Question   dbQuestion `gorm:"foreignKey:QuestionID;references:ID"`
}

func CreateFormsQuestion(f *domain.Form, q *domain.Question, db *gorm.DB) error {
	dbQ := dbFormsQuestion{
		ID:         uuid.NewString(),
		FormID:     f.ID,
		Form:       dbForm{},
		QuestionID: q.ID,
		Question:   dbQuestion{},
	}

	err := db.Create(&dbQ).Error
	if err != nil {
		return err
	}
	return db.Save(&dbQ).Error
}
