package database

import (
	"tusur-forms/internal/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type dbPossibleAnswer struct {
	ID      string `gorm:"primaryKey"`
	Content string
}

type dbQuestionPossibleAnswer struct {
	ID               string `gorm:"primaryKey"`
	QuestionID       string
	Question         dbQuestion `gorm:"foreignKey:QuestionID;references:ID"`
	PossibleAnswerID string
	PossibleAnswer   dbPossibleAnswer `gorm:"foreignKey:PossibleAnswerID;references:ID"`
}

func CreatePossibleAnswer(pa *domain.PossibleAnswer, q *domain.Question, db *gorm.DB) error {
	dbPa := dbPossibleAnswer{
		ID:      uuid.NewString(),
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
		ID:               uuid.NewString(),
		QuestionID:       q.ID,
		Question:         dbQuestion{},
		PossibleAnswerID: pa.ID,
		PossibleAnswer:   dbPossibleAnswer{},
	}

	err := db.Create(&dbF).Error
	if err != nil {
		return err
	}
	return db.Save(&dbF).Error
}
