package repository

import (
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
	dbQID, err := g.GetQuestionIDByTitleAndDesc(q.Title, q.Description)
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
		for _, item := range q.PossibleAnswers {
			pa, err := g.getPossibleAnswer(item)
			if err != nil {
				return "", err
			}
			if pa == nil {
				_, err = g.CreatePossibleAnswer(item, qID)
				if err != nil {
					return "", err
				}
			} else {
				err = g.CreateQuestionPossibleAnswer(pa, qID)
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

func (g *GormRepository) GetQuestionIDByTitleAndDesc(qTitle string, qDesc string) (string, error) {
	var dbQ []*dbQuestion
	err := g.db.Where("title = ? AND description = ?", qTitle, qDesc).Limit(1).Find(&dbQ).Error
	if err != nil {
		return "", err
	}
	if len(dbQ) == 0 {
		return "", nil
	}
	return dbQ[0].ID, nil
}

func (g *GormRepository) GetQuestionByTitle(qTitle string) (*domain.Question, error) {
	var dbQ []*dbQuestion
	err := g.db.
		Preload("QuestionType").
		Preload("QuestionPossibleAnswers.PossibleAnswer").
		Where("title = ?", qTitle).Limit(1).Find(&dbQ).Error
	if err != nil {
		return nil, err
	}
	if len(dbQ) == 0 {
		return nil, nil
	}
	pa := make([]*domain.PossibleAnswer, 0, len(dbQ[0].QuestionPossibleAnswers))
	for _, p := range dbQ[0].QuestionPossibleAnswers {
		pa = append(pa, &domain.PossibleAnswer{
			Content: p.PossibleAnswer.Content,
		})
	}
	return &domain.Question{
		Title:       dbQ[0].Title,
		Description: dbQ[0].Description,
		Type: domain.QuestionType{
			Title: domain.QuestionTypeTitles(dbQ[0].QuestionType.Title),
		},
		IsRequired:      dbQ[0].IsRequired,
		Answers:         nil,
		PossibleAnswers: pa,
	}, nil
}
func (g *GormRepository) GetQuestions() ([]*domain.Question, error) {
	var questions []*dbQuestion
	err := g.db.
		Preload("QuestionType").
		Preload("QuestionPossibleAnswers.PossibleAnswer").
		Find(&questions).Error
	if err != nil {
		return nil, err
	}
	qs := make([]*domain.Question, 0, len(questions))
	for _, q := range questions {
		pa := make([]*domain.PossibleAnswer, 0, len(q.QuestionPossibleAnswers))
		for _, p := range q.QuestionPossibleAnswers {
			pa = append(pa, &domain.PossibleAnswer{
				Content: p.PossibleAnswer.Content,
			})
		}
		qs = append(qs, &domain.Question{
			Title:       q.Title,
			Description: q.Description,
			Type: domain.QuestionType{
				Title: domain.QuestionTypeTitles(q.QuestionType.Title),
			},
			IsRequired:      q.IsRequired,
			PossibleAnswers: pa,
		})
	}
	return qs, nil
}
