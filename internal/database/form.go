package database

import (
	"errors"
	"time"
	"tusur-forms/internal/domain"

	"gorm.io/gorm"
)

type dbForm struct {
	ID             string `gorm:"primaryKey"`
	Title          string
	DocumentTitle  string
	ExternalID     string
	CreatedAt      time.Time
	FormsQuestions []*dbFormsQuestion `gorm:"foreignKey:FormID;references:ID"`
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
	var form dbForm
	//var dbFormQuestions []*dbFormsQuestion
	var domainQuestions []*domain.Question

	err := g.db.Preload("FormsQuestions.Question.QuestionType").
		Preload("FormsQuestions.Question.QuestionPossibleAnswers.PossibleAnswer").
		Where("id = ?", internalID).
		First(&form).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	for _, item := range form.FormsQuestions {
		q := item.Question
		domainPossibleAnswers := make([]*domain.PossibleAnswer, 0, len(q.QuestionPossibleAnswers))

		for _, inItem := range q.QuestionPossibleAnswers {
			p := inItem.PossibleAnswer
			domainPossibleAnswers = append(domainPossibleAnswers, &domain.PossibleAnswer{
				Content: p.Content,
			})
		}
		domainQuestions = append(domainQuestions, &domain.Question{
			Title:       q.Title,
			Description: q.Description,
			Type: domain.QuestionType{
				Title: domain.QuestionTypeTitles(q.QuestionType.Title),
			},
			IsRequired:      q.IsRequired,
			Answers:         []*domain.Answer{}, //TODO add getting
			PossibleAnswers: domainPossibleAnswers,
		})
	}
	resultDomain := domain.Form{
		ID:            internalID,
		ExternalID:    form.ExternalID,
		Title:         form.Title,
		DocumentTitle: form.DocumentTitle,
		CreatedAt:     form.CreatedAt,
		Questions:     domainQuestions,
	}
	return &resultDomain, nil
}
