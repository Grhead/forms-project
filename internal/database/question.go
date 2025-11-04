package database

import (
	"tusur-forms/internal/domain"

	"gorm.io/gorm"
)

type dbQuestion struct {
	Id           string `gorm:"primaryKey"`
	Title        string
	Description  string
	TypeId       string
	QuestionType dbQuestionType `gorm:"foreignKey:TypeId;references:Id"`
	IsRequired   bool
}

type dbQuestionType struct {
	Id    string `gorm:"primaryKey"`
	Title string
}

func CreateQuestion(q *domain.Question, db *gorm.DB) error {
	dbQ := dbQuestion{
		Id:          q.Id,
		Title:       q.Title,
		Description: q.Description,
		IsRequired:  false,
		TypeId:      q.Type.Id,
	}

	err := db.Create(&dbQ).Error
	if err != nil {
		return err
	}
	return db.Save(&dbQ).Error
}

func CreateQuestionType(qt *domain.QuestionType, db *gorm.DB) error {
	dbQt := dbQuestionType{
		Id:    qt.Id,
		Title: string(qt.Title),
	}

	err := db.Create(&dbQt).Error
	if err != nil {
		return err
	}
	return db.Save(&dbQt).Error
}
