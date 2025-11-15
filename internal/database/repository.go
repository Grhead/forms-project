package database

import (
	"tusur-forms/internal/domain"

	"gorm.io/gorm"
)

type GormRepository struct {
	db *gorm.DB
}

type FormRepository interface {
	CreateForm(f *domain.Form) error
	CreateQuestion(q *domain.Question) error
	CreateQuestionType(qt *domain.QuestionType) error
	CreatePossibleAnswer(pa *domain.PossibleAnswer, q *domain.Question) (*dbPossibleAnswer, error)
	createQuestionPossibleAnswer(pa *dbPossibleAnswer, q *domain.Question) error
	CreateFormsQuestion(f *domain.Form, q *domain.Question) error
	CreateAnswer(a *domain.Answer) error

	getFormsQuestionID(a *domain.Answer) (dbFormsQuestion, error)
	GetForm(internalID string) (*domain.Form, error)
	GetFormExternalID(internalID string) (string, error)

	Migrate() error
	CheckExists() (bool, error)
}

func NewGormRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{
		db: db,
	}
}
