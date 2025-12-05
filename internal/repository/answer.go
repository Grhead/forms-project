package repository

import (
	"time"
	"tusur-forms/internal/domain"

	"github.com/google/uuid"
)

type dbAnswer struct {
	ID                  string `gorm:"primaryKey"`
	ResponseEnvironment string
	SubmittedAt         time.Time
	Content             string
	FormsQuestionID     string
	FormsQuestion       *dbFormsQuestions `gorm:"foreignKey:FormsQuestionID;references:ID"`
}

func (g *GormRepository) CreateAnswer(a *domain.Answer,
	formID string,
	questionID string,
	responseEnvironment string) error {
	formsQuestionID, err := g.getFormsQuestionID(formID, questionID)
	if err != nil {
		return err
	}
	exists, err := g.checkResponseEnvironmentExists(responseEnvironment, formsQuestionID)
	if err != nil {
		return err
	}

	if exists {
		return nil
	}
	dbQ := dbAnswer{
		ID:                  uuid.NewString(),
		ResponseEnvironment: responseEnvironment,
		SubmittedAt:         a.SubmittedAt,
		Content:             a.Content,
		FormsQuestionID:     formsQuestionID,
	}
	err = g.db.Create(&dbQ).Error
	if err != nil {
		return err
	}
	return nil
}

func (g *GormRepository) getFormsQuestionID(formID string, questionID string) (string, error) {
	var fq []*dbFormsQuestions
	err := g.db.
		Where("form_id = ? AND question_id = ?", formID, questionID).
		Limit(1).
		Find(&fq).Error
	if err != nil {
		return "", err
	}
	if len(fq) == 0 {
		return "", nil
	}
	return fq[0].ID, nil
}

func (g *GormRepository) checkResponseEnvironmentExists(environment string, formsQuestionID string) (bool, error) {
	var dbRespEnv string
	err := g.db.
		Table("db_answers").
		Where("response_environment = ? AND forms_question_id = ?", environment, formsQuestionID).
		Select("id").
		Find(&dbRespEnv).Error
	if err != nil {
		return false, err
	}
	if len(dbRespEnv) == 0 {
		return false, nil
	}
	return true, nil
}

func (g *GormRepository) GetAnswers(formID string, questionID string) ([]*domain.Answer, error) {
	var dbAnswers []*dbAnswer
	err := g.db.
		Joins("FormsQuestion").
		Where("FormsQuestion.form_id = ? AND FormsQuestion.question_id = ?", formID, questionID).
		Preload("FormsQuestion").
		Find(&dbAnswers).
		Error
	if err != nil {
		return nil, err
	}
	answers := make([]*domain.Answer, 0, len(dbAnswers))
	for _, dbA := range dbAnswers {
		ans := domain.Answer{
			SubmittedAt: dbA.SubmittedAt,
			Content:     dbA.Content,
		}
		answers = append(answers, &ans)
	}
	return answers, nil
}
