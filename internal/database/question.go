package database

import (
	"errors"
	"tusur-forms/internal/domain"

	"gorm.io/gorm"
)

type dbQuestion struct {
	ID                      string `gorm:"primaryKey"`
	Title                   string
	Description             string
	TypeID                  string
	QuestionType            dbQuestionType `gorm:"foreignKey:TypeID;references:ID"`
	IsRequired              bool
	QuestionPossibleAnswers []dbQuestionPossibleAnswer `gorm:"foreignKey:QuestionID;references:ID"`
}

type dbQuestionType struct {
	ID    string `gorm:"primaryKey"`
	Title string
}

func (g *GormRepository) CreateQuestion(q *domain.Question) error {
	dbQt, err := g.getQuestionTypeByID(q.Type.ID)
	if err != nil {
		return err
	}
	if dbQt == nil {
		err = g.createQuestionType(&q.Type)
		if err != nil {
			return err
		}
	}
	dbQ := dbQuestion{
		ID:          q.ID,
		Title:       q.Title,
		Description: q.Description,
		IsRequired:  q.IsRequired,
		TypeID:      q.Type.ID,
	}

	err = g.db.Create(&dbQ).Error
	if err != nil {
		return err
	}
	if q.Type.Title == domain.TypeCheckbox || q.Type.Title == domain.TypeRadio {
		for _, item := range q.PossibleAnswers {
			paID, err := g.getPossibleAnswer(item)
			if err != nil {
				return err
			}
			if paID == nil {
				_, err = g.CreatePossibleAnswer(item, q)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (g *GormRepository) createQuestionType(qt *domain.QuestionType) error {
	dbQt := dbQuestionType{
		ID:    qt.ID,
		Title: string(qt.Title),
	}

	err := g.db.Create(&dbQt).Error
	if err != nil {
		return err
	}
	return nil
}

func (g *GormRepository) getQuestionTypeByID(qtID string) (*dbQuestionType, error) {
	var dbQt dbQuestionType
	err := g.db.Where("id = ?", qtID).First(&dbQt).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &dbQt, nil
}
