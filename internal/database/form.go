package database

import (
	"errors"
	"time"
	"tusur-forms/internal/domain"

	"gorm.io/gorm"
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
	return nil
}
func (g *GormRepository) GetFormExternalID(internalID string) (string, error) {
	var form []*dbForm

	err := g.db.Where("id = ?", internalID).Limit(1).Select("external_id").Find(&form).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "", nil
	}
	if err != nil {
		return "", err
	}
	return form[0].ExternalID, nil
}

func (g *GormRepository) GetForm(internalID string) (*domain.Form, error) {
	var dbForm dbForm
	var dbFormQuestions []*dbFormsQuestion
	var domainQuestions []*domain.Question

	err := g.db.Where("id = ?", internalID).First(&dbForm).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	err = g.db.Where("form_id = ?", internalID).
		Preload("Question.QuestionType").
		Preload("Question.QuestionPossibleAnswers.PossibleAnswer").
		Find(&dbFormQuestions).Error
	if err != nil {
		return nil, err
	}
	for _, item := range dbFormQuestions {
		q := item.Question
		domainPossibleAnswers := make([]*domain.PossibleAnswer, 0, len(q.QuestionPossibleAnswers))

		for _, inItem := range q.QuestionPossibleAnswers {
			p := inItem.PossibleAnswer
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
		ExternalID:    dbForm.ExternalID,
		Title:         dbForm.Title,
		DocumentTitle: dbForm.DocumentTitle,
		CreatedAt:     dbForm.CreatedAt,
		Questions:     domainQuestions,
	}
	return &resultDomain, nil
}
