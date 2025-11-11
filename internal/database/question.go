package database

import (
	"log"
	"tusur-forms/internal/domain"
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

func (g *GormRepository) CreateQuestion(q *domain.Question) error {
	dbQ := dbQuestion{
		ID:          q.ID,
		Title:       q.Title,
		Description: q.Description,
		IsRequired:  false,
		TypeID:      q.Type.ID,
	}

	err := g.db.Create(&dbQ).Error
	if err != nil {
		return err
	}
	if q.Type.Title == domain.TypeCheckbox || q.Type.Title == domain.TypeRadio {
		for _, item := range q.PossibleAnswers {
			paID, err := g.getPossibleAnswer(&item)
			if err != nil {
					return err
			}
			if paID == nil {
				_, err = g.CreatePossibleAnswer(&item, q)
				if err != nil {
					return err
				}
			}
		}
	}

	return g.db.Save(&dbQ).Error
}

func (g *GormRepository) CreateQuestionType(qt *domain.QuestionType) error {
	dbQt := dbQuestionType{
		ID:    qt.ID,
		Title: string(qt.Title),
	}

	err := g.db.Create(&dbQt).Error
	if err != nil {
		return err
	}
	return g.db.Save(&dbQt).Error
}
