package database

import (
	"tusur-forms/internal/domain"

	"gorm.io/gorm"
)

type dbQuestion struct {
	ID           string `gorm:"primaryKey"`
	Title        string
	Description  string
	TypeID       string
	QuestionType dbQuestionType `gorm:"foreignKey:TypeID;references:ID"`
	IsRequired   bool
}

type dbQuestionType struct {
	ID    string `gorm:"primaryKey"`
	Title string
}

func CreateQuestion(q *domain.Question, db *gorm.DB) error {
	dbQ := dbQuestion{
		ID:          q.ID,
		Title:       q.Title,
		Description: q.Description,
		IsRequired:  false,
		TypeID:      q.Type.ID,
	}

	err := db.Create(&dbQ).Error
	if err != nil {
		return err
	}
	return db.Save(&dbQ).Error
}

func CreateQuestionType(qt *domain.QuestionType, db *gorm.DB) error {
	dbQt := dbQuestionType{
		ID:    qt.ID,
		Title: string(qt.Title),
	}

	err := db.Create(&dbQt).Error
	if err != nil {
		return err
	}
	return db.Save(&dbQt).Error
}
