package repository

import (
	"log"
	"tusur-forms/internal/domain"

	"github.com/google/uuid"
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
	dbQID, err := g.GetQuestionIDByTitle(q.Title)
	if err != nil {
		return "", err
	}
	if dbQID == "" { //TODO
		qID = uuid.NewString()
		err = g.db.Create(&dbQuestion{
			ID:          qID,
			Title:       q.Title,
			Description: q.Description,
			IsRequired:  q.IsRequired,
			TypeID:      dbQt.ID,
		}).Error
		if err != nil {
			return "", err
		}
	} else {
		qID = dbQID
	}

	if q.Type.Title == domain.TypeCheckbox || q.Type.Title == domain.TypeRadio {
		log.Println("Im creation")
		for _, item := range q.PossibleAnswers {
			log.Println(item.Print())

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
	log.Println("Create Question Type")
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
	var dbQt []*dbQuestionType
	err := g.db.Where("title = ?", qtTitle).Limit(1).Find(&dbQt).Error
	if err != nil {
		return nil, err
	}
	if len(dbQt) == 0 {
		return nil, nil
	}
	return dbQt[0], nil
}

func (g *GormRepository) GetQuestionIDByTitle(qTitle string) (string, error) {
	var dbQ []*dbQuestion
	err := g.db.Where("title = ?", qTitle).Limit(1).Find(&dbQ).Error
	if err != nil {
		return "", err
	}
	if len(dbQ) == 0 {
		return "", nil
	}
	return dbQ[0].ID, nil
}

func (g *GormRepository) GetQuestionIDs(formID string) ([]string, error) {
	var questions []string
	err := g.db.Table("db_forms_questions").Where("form_id = ?", formID).Select("question_id").Find(&questions).Error
	if err != nil {
		return nil, err
	}
	return questions, nil
}
