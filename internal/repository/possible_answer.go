package repository

import (
	"tusur-forms/internal/domain"

	"github.com/google/uuid"
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

func (g *GormRepository) CreatePossibleAnswer(pa *domain.PossibleAnswer, qID string) (*domain.PossibleAnswer, error) {
	dbPa := dbPossibleAnswer{
		ID:      uuid.NewString(),
		Content: pa.Content,
	}
	err := g.createQuestionPossibleAnswer(&dbPa, qID)
	if err != nil {
		return nil, err
	}
	err = g.db.Create(&dbPa).Error
	if err != nil {
		return nil, err
	}
	return &domain.PossibleAnswer{Content: dbPa.Content}, nil
}

func (g *GormRepository) createQuestionPossibleAnswer(pa *dbPossibleAnswer, qID string) error {
	dbF := dbQuestionPossibleAnswer{
		ID:               uuid.NewString(),
		QuestionID:       qID,
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
	err := g.db.Where("content = ?", a.Content).Limit(1).Find(&fq).Error
	if err != nil {
		return nil, err
	}
	if len(fq) == 0 {
		return nil, nil
	}
	return fq[0], nil
}
