package database

import (
	"tusur-forms/internal/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type dbPossibleAnswer struct {
	Id      string `gorm:"primaryKey"`
	Content string
}

type dbQuestionPossibleAnswer struct {
	Id               string `gorm:"primaryKey"`
	QuestionId       string
	Question         dbQuestion `gorm:"foreignKey:QuestionId;references:Id"`
	PossibleAnswerId string
	PossibleAnswer   dbPossibleAnswer `gorm:"foreignKey:PossibleAnswerId;references:Id"`
}

func CreatePossibleAnswer(pa *domain.PossibleAnswer, q *domain.Question, db *gorm.DB) error {
	dbPa := dbPossibleAnswer{
		Id:      pa.Id,
		Content: pa.Content,
	}
	err := createQuestionPossibleAnswer(&dbPa, q, db)
	if err != nil {
		return err
	}
	err = db.Create(&dbPa).Error
	if err != nil {
		return err
	}
	return db.Save(&dbPa).Error
}

func createQuestionPossibleAnswer(pa *dbPossibleAnswer, q *domain.Question, db *gorm.DB) error {
	dbF := dbQuestionPossibleAnswer{
		Id:               uuid.NewString(),
		QuestionId:       q.Id,
		Question:         dbQuestion{},
		PossibleAnswerId: pa.Id,
		PossibleAnswer:   dbPossibleAnswer{},
	}

	err := db.Create(&dbF).Error
	if err != nil {
		return err
	}
	return db.Save(&dbF).Error
}
