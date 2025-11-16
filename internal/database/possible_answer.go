package database

import (
	"errors"
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

func (g *GormRepository) CreatePossibleAnswer(pa *domain.PossibleAnswer, q *domain.Question) (*domain.PossibleAnswer, error) {
	dbPa := dbPossibleAnswer{
		ID:      uuid.NewString(),
		Content: pa.Content,
	}
	err := g.createQuestionPossibleAnswer(&dbPa, q)
	if err != nil {
		return nil, err
	}
	err = g.db.Create(&dbPa).Error
	if err != nil {
		return nil, err
	}
	return &domain.PossibleAnswer{Content: dbPa.Content}, nil
}

func (g *GormRepository) createQuestionPossibleAnswer(pa *dbPossibleAnswer, q *domain.Question) error {
	dbF := dbQuestionPossibleAnswer{
		ID:               uuid.NewString(),
		QuestionID:       q.ID,
		Question:         dbQuestion{},
		PossibleAnswerID: pa.ID,
		PossibleAnswer:   dbPossibleAnswer{},
	}

	err := g.db.Create(&dbF).Error
	if err != nil {
		return err
	}
	return nil
}

func (g *GormRepository) getPossibleAnswer(a *domain.PossibleAnswer) (*dbPossibleAnswer, error) {
	var fq []*dbPossibleAnswer
	err := g.db.Where("content = ?", a.Content).Find(&fq).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return fq[0], nil
}
