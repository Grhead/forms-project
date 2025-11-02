package database

import (
	"tusur-forms/internal/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type dbFormsQuestion struct {
	Id         string `gorm:"primaryKey"`
	FormId     string
	Form       dbForm `gorm:"foreignKey:FormId;references:Id"`
	QuestionId string
	Question   dbQuestion `gorm:"foreignKey:QuestionId;references:Id"`
}

func CreateFormsQuestion(f *domain.Form, q *domain.Question, db *gorm.DB) error {
	dbQ := dbFormsQuestion{
		Id:         uuid.NewString(),
		FormId:     f.Id,
		Form:       dbForm{},
		QuestionId: q.Id,
		Question:   dbQuestion{},
	}

	err := db.Create(&dbQ).Error
	if err != nil {
		return err
	}
	return db.Save(&dbQ).Error
}
