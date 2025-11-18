package database

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
	FormsQuestion       dbFormsQuestion `gorm:"foreignKey:FormsQuestionID;references:ID"`
}

func (g *GormRepository) CreateAnswer(a *domain.Answer, formID string, questionID string, responseEnvironment string) error {
	fqID, err := g.getFormsQuestionID(formID, questionID)
	if err != nil {
		return err
	}
	dbQ := dbAnswer{
		ID:                  uuid.NewString(),
		ResponseEnvironment: responseEnvironment,
		SubmittedAt:         a.SubmittedAt,
		Content:             a.Content,
		FormsQuestionID:     fqID,
		FormsQuestion:       dbFormsQuestion{},
	}
	err = g.db.Create(&dbQ).Error
	if err != nil {
		return err
	}
	return nil
}

func (g *GormRepository) getFormsQuestionID(formID string, questionID string) (string, error) {
	var fq dbFormsQuestion
	err := g.db.Where("form_id = ? AND question_id = ?", formID, questionID).First(&fq).Error
	if err != nil {
		return "", err
	}
	return fq.ID, nil
}

func (g *GormRepository) CheckResponseEnvironmentExists(environment string) (bool, error) {
	var dbRespEnv string
	err := g.db.Table("db_answers").Where("response_environment = ?", environment).Select("id").Find(&dbRespEnv).Error
	if err != nil {
		return false, err
	}
	if len(dbRespEnv) == 0 {
		return false, nil
	}
	return true, nil
}
