package orchectrators

import (
	"tusur-forms/internal/database"
	"tusur-forms/internal/domain"
	"tusur-forms/internal/services/google"
)

type FormsOrchestrator struct {
	creator    google.FormService
	repository database.FormRepository
}

func NewFormsOrchestrator(c google.FormService, r database.FormRepository) *FormsOrchestrator {
	return &FormsOrchestrator{
		creator:    c,
		repository: r,
	}
}

func (s *FormsOrchestrator) CheckoutForm(title string, documentTitle string, questions ...[]*domain.Question) (domain.Form, error) {
	d, err := s.creator.NewForm(title, documentTitle)
	if err != nil {
		return domain.Form{}, err
	}
	s.repository.CreateForm(&d)
	if len(questions) != 0 {
		d, err := s.creator.SetQuestions(d, questions[0])
		if err != nil {
			return domain.Form{}, err
		}
		s.repository.CreateQuestion(questions[0][0])
		s.repository.CreateFormsQuestion(&d, questions[0][0])
		return d, nil
	}
	return d, nil
}
