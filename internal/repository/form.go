package repository

import (
	"time"
	"tusur-forms/internal/domain"
)

type dbForm struct {
	ID             string `gorm:"primaryKey"`
	Title          string
	DocumentTitle  string
	ExternalID     string
	CreatedAt      time.Time
	FormsQuestions []*dbFormsQuestions `gorm:"foreignKey:FormID;references:ID"`
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
	err := g.db.
		Where("id = ?", internalID).
		Limit(1).
		Select("external_id").
		Find(&form).Error
	if err != nil {
		return "", err
	}
	if len(form) == 0 {
		return "", nil
	}
	return form[0].ExternalID, nil
}

func (g *GormRepository) GetForm(internalID string) (*domain.Form, error) {
	var forms []*dbForm
	var domainQuestions []*domain.Question

	err := g.db.
		Preload("FormsQuestions.Question.QuestionType").
		Preload("FormsQuestions.Question.QuestionPossibleAnswers.PossibleAnswer").
		Where("id = ?", internalID).
		First(&forms).Error
	if err != nil {
		return nil, err
	}
	if len(forms) == 0 {
		return nil, nil
	}
	var form = forms[0]
	for _, item := range form.FormsQuestions {
		q := item.Question
		domainPossibleAnswers := make([]*domain.PossibleAnswer, 0, len(q.QuestionPossibleAnswers))
		for _, inItem := range q.QuestionPossibleAnswers {
			p := inItem.PossibleAnswer
			domainPossibleAnswers = append(domainPossibleAnswers, &domain.PossibleAnswer{
				Content: p.Content,
			})
		}
		answers, err := g.GetAnswers(internalID, q.ID)
		if err != nil {
			return nil, err
		}
		domainQuestions = append(domainQuestions, &domain.Question{
			Title:       q.Title,
			Description: q.Description,
			Type: domain.QuestionType{
				Title: domain.QuestionTypeTitles(q.QuestionType.Title),
			},
			IsRequired:      q.IsRequired,
			Answers:         answers,
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
