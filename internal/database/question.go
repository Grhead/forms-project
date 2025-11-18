package database

import (
	"errors"
	"tusur-forms/internal/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type dbQuestion struct {
	ID                      string `gorm:"primaryKey"`
	Title                   string
	Description             string
	TypeID                  string
	QuestionType            dbQuestionType `gorm:"foreignKey:TypeID;references:ID"`
	IsRequired              bool
	QuestionPossibleAnswers []*dbQuestionPossibleAnswer `gorm:"foreignKey:QuestionID;references:ID"`
}

type dbQuestionType struct {
	ID    string `gorm:"primaryKey"`
	Title string
}

func (g *GormRepository) CreateQuestion(q *domain.Question) (string, error) {
	var qID string
	dbQt, err := g.getQuestionTypeByTitle(string(q.Type.Title))
	if err != nil {
		return "", err
	}
	if dbQt == nil {
		dbQt, err = g.createQuestionType(&q.Type)
		if err != nil {
			return "", err
		}
	}
	dbQ, err := g.GetQuestionByTitle(string(q.Title))
	if err != nil {
		return "", err
	}
	if dbQ == nil {
		qID = uuid.NewString()
		dbQ = &dbQuestion{
			ID:          qID,
			Title:       q.Title,
			Description: q.Description,
			IsRequired:  q.IsRequired,
			TypeID:      dbQt.ID,
		}
		err = g.db.Create(&dbQ).Error
		if err != nil {
			return "", err
		}
	} else {
		qID = dbQ.ID
	}

	if q.Type.Title == domain.TypeCheckbox || q.Type.Title == domain.TypeRadio {
		for _, item := range q.PossibleAnswers {
			paID, err := g.getPossibleAnswer(item)
			if err != nil {
				return "", err
			}
			if paID == nil {
				_, err = g.CreatePossibleAnswer(item, qID)
				if err != nil {
					return "", err
				}
			}
		}
	}

	return qID, nil
}

func (g *GormRepository) createQuestionType(qt *domain.QuestionType) (*dbQuestionType, error) {
	dbQt := dbQuestionType{
		ID:    uuid.NewString(),
		Title: string(qt.Title),
	}
	err := g.db.Create(&dbQt).Error
	if err != nil {
		return nil, err
	}
	return &dbQt, nil
}

func (g *GormRepository) getQuestionTypeByTitle(qtTitle string) (*dbQuestionType, error) {
	var dbQt dbQuestionType
	err := g.db.Where("title = ?", qtTitle).First(&dbQt).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &dbQt, nil
}

func (g *GormRepository) GetQuestionByTitle(qTitle string) (*dbQuestion, error) {
	var dbQ dbQuestion
	err := g.db.Where("title = ?", qTitle).First(&dbQ).Error
	if err != nil {
		return nil, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &dbQ, nil
}

func (g *GormRepository) GetQuestionIDs(formID string) ([]string, error) {
	var questions []string
	err := g.db.Table("db_forms_questions").Where("form_id = ?", formID).Select("question_id").Find(&questions).Error
	if err != nil {
		return nil, err
	}
	return questions, nil
}
