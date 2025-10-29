package database

import (
	"tusur-forms/internal/domain"

	"gorm.io/gorm"
)

type dbQuestion struct {
	QuestionId     string `gorm:"primaryKey"`
	QuestionTitle  string
	QuestionTypeId string
	QuestionType   dbQuestionType `gorm:"references:QuestionTypeId"`
	IsRequired     bool
}

type dbQuestionType struct {
	QuestionTypeId    string `gorm:"primaryKey"`
	QuestionTypeTitle string
}

func CreateQuestion(question *domain.Question, db *gorm.DB) {
	dbquestion := dbQuestion{
		QuestionId:     question.Id,
		QuestionTitle:  question.Title,
		IsRequired:     false,
		QuestionTypeId: question.Type.Id,
	}

	db.Create(&dbquestion)
	db.Save(&dbquestion)
}

func CreateQuestionType(qtype *domain.QuestionType, db *gorm.DB) {
	dbtype := dbQuestionType{
		QuestionTypeId:    qtype.Id,
		QuestionTypeTitle: qtype.Title,
	}

	db.Create(&dbtype)
	db.Save(&dbtype)
}
