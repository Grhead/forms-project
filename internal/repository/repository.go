package repository

import (
	"tusur-forms/internal/domain"

	"gorm.io/gorm"
)

type GormRepository struct {
	db *gorm.DB
}

type FormRepository interface {
	CreateForm(f *domain.Form) error
	CreateQuestion(q *domain.Question) (string, error)
	CreateAnswer(a *domain.Answer, formID string, questionID string, environmentID string) error
	CreateFormsQuestion(fID string, qID string) error
	createQuestionType(qt *domain.QuestionType) (*dbQuestionType, error)
	CreatePossibleAnswer(pa *domain.PossibleAnswer, qID string) (*domain.PossibleAnswer, error)
	createQuestionPossibleAnswer(pa *dbPossibleAnswer, qID string) error

	GetForm(internalID string, isExternal bool) (*domain.Form, error)
	GetForms() ([]*domain.Form, error)
	GetFormExternalID(internalID string) (string, error)
	GetAnswers(formID string, questionID string) ([]*domain.Answer, error)
	GetQuestionIDs(formID string) ([]string, error)
	GetQuestions() ([]*domain.Question, error)
	GetQuestionIDByTitle(qTitle string) (string, error)
	getQuestionTypeByTitle(qtID string) (*dbQuestionType, error)
	getFormsQuestionID(formID string, questinID string) (string, error)

	CheckResponseEnvironmentExists(environment string) (bool, error)
	Migrate() error
	CheckExists() (bool, error)
}

func NewGormRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{
		db: db,
	}
}
