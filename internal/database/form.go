package database

import (
	"log"
	"time"
	"tusur-forms/internal/domain"
)

type dbForm struct {
	ID            string `gorm:"primaryKey"`
	Title         string
	DocumentTitle string
	ExternalID    string
	CreatedAt     time.Time
}

func (g *GormRepository) CreateForm(f *domain.Form) error {
	dbF := dbForm{
		ID:            f.ID,
		ExternalID:    f.ExternalID,
		CreatedAt:     f.CreatedAt,
		Title:         f.Title,
		DocumentTitle: f.DocumentTitle,
	}

	err := g.db.Create(&dbF).Error
	if err != nil {
		return err
	}
	return g.db.Save(&dbF).Error
}
func (g *GormRepository) GetFormExternalID(internalID string) (string, error) {
	var form []*dbForm

	err := g.db.Where("id = ?", internalID).Limit(1).Select("external_id").Find(&form).Error
	if err != nil {
		log.Println(err)
		return "", err
	} else if len(form) == 0 {
		return "", nil
	}
	return form[0].ExternalID, nil
}

func (g *GormRepository) GetForm(internalID string) (*domain.Form, error) {
	var dbForm []*dbForm
	var dbFormQuestions []*dbFormsQuestion
	var domainQuestions []*domain.Question

	err := g.db.Where("id = ?", internalID).Limit(1).Find(&dbForm).Error
	if err != nil {
		return nil, err
	} else if len(dbForm) == 0 {
		return nil, nil
	}
	err = g.db.Where("form_id = ?", internalID).Preload("Question.QuestionType").Find(&dbFormQuestions).Error

	if err != nil {
		return nil, err
	} else if len(dbFormQuestions) == 0 {
		return nil, nil
	}
	for i := range dbFormQuestions {
		var dbQuestionPossibleAnswers []*dbQuestionPossibleAnswer
		var domainPossibleAnswers []*domain.PossibleAnswer

		q := dbFormQuestions[i].Question
		err = g.db.Where("question_id = ?", q.ID).Preload("PossibleAnswer").Find(&dbQuestionPossibleAnswers).Error
		if err != nil {
			return nil, err
		} else if len(dbFormQuestions) == 0 {
			return nil, nil
		}
		log.Println(dbQuestionPossibleAnswers)

		for j := range dbQuestionPossibleAnswers {
			p := dbQuestionPossibleAnswers[j].PossibleAnswer

			domainPossibleAnswers = append(domainPossibleAnswers, &domain.PossibleAnswer{
				Content: p.Content,
			})
		}
		domainQuestions = append(domainQuestions, &domain.Question{
			ID:          q.ID,
			Title:       q.Title,
			Description: q.Description,
			Type: domain.QuestionType{
				ID:    q.QuestionType.ID,
				Title: domain.QuestionTypeTitles(q.QuestionType.Title),
			},
			IsRequired:      q.IsRequired,
			Answers:         []*domain.Answer{}, //TODO add getting
			PossibleAnswers: domainPossibleAnswers,
		})
	}
	resultDomain := domain.Form{
		ID:            internalID,
		ExternalID:    dbForm[0].ExternalID,
		Title:         dbForm[0].Title,
		DocumentTitle: dbForm[0].DocumentTitle,
		CreatedAt:     dbForm[0].CreatedAt,
		Questions:     domainQuestions,
	}
	return &resultDomain, nil
}
