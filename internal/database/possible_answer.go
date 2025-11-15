package database

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

func (g *GormRepository) CreatePossibleAnswer(pa *domain.PossibleAnswer, q *domain.Question) (*dbPossibleAnswer, error) {
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
	return &dbPa, g.db.Save(&dbPa).Error
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
	return g.db.Save(&dbF).Error
}

func (g *GormRepository) getPossibleAnswer(a *domain.PossibleAnswer) (*dbPossibleAnswer, error) {
	var fq []*dbPossibleAnswer
	err := g.db.Where("content = ?", a.Content).Find(&fq)
	if err != nil {
		return nil, err.Error
	} else if len(fq) == 0 {
		return nil, nil
	}
	return	fq[0], nil
}
