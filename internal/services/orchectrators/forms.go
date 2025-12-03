package orchectrators

import (
	"tusur-forms/internal/domain"
	"tusur-forms/internal/repository"
	service "tusur-forms/internal/services/forms"
)

type FormsOrchestrator struct {
	creator    service.FormService
	repository repository.FormRepository
}

func NewFormsOrchestrator(c service.FormService, r repository.FormRepository) *FormsOrchestrator {
	return &FormsOrchestrator{
		creator:    c,
		repository: r,
	}
}

func (s *FormsOrchestrator) CheckoutForm(title string, documentTitle string, questions ...[]*domain.Question) (*domain.Form, error) {
	var form *domain.Form
	d, err := s.creator.NewForm(title, documentTitle)
	if err != nil {
		return nil, err
	}
	err = s.repository.CreateForm(&d)
	if err != nil {
		return nil, err
	}
	if len(questions) != 0 {
		err = s.creator.SetQuestions(&d, questions[0])
		if err != nil {
			return nil, err
		}
		for i := range questions[0] {
			qID, err := s.repository.CreateQuestion(questions[0][i])
			if err != nil {
				return nil, err
			}
			err = s.repository.CreateFormsQuestion(d.ID, qID)
			if err != nil {
				return nil, err
			}
		}
		return form, nil
	}
	form, err = s.repository.GetForm(d.ID, false)
	return form, nil
}

func (s *FormsOrchestrator) CheckoutAnswers(formID string) (*domain.Form, error) {
	form, err := s.creator.GetForm(formID)
	if err != nil {
		return nil, err
	}
	for _, item := range form.Questions {
		for _, f := range item.Answers {
			exists, err := s.repository.CheckResponseEnvironmentExists(f.ResponseID)
			if err != nil {
				return nil, err
			}
			if exists {
				continue
			}
			err = s.repository.CreateAnswer(f.ToDomain(), formID, item.ID, f.ResponseID)
			if err != nil {
				return nil, err
			}
		}
	}
	domainForm, err := s.repository.GetForm(formID, false)
	if err != nil {
		return nil, err
	}
	return domainForm, nil
}

func (s *FormsOrchestrator) GetForm(ID string, isExternal bool) (*domain.Form, error) {
	forms, err := s.repository.GetForm(ID, isExternal)
	if err != nil {
		return nil, err
	}
	return forms, nil
}

func (s *FormsOrchestrator) GetForms() ([]*domain.Form, error) {
	forms, err := s.repository.GetForms()
	if err != nil {
		return nil, err
	}
	return forms, nil
}

func (s *FormsOrchestrator) GetQuestions() ([]*domain.Question, error) {
	questions, err := s.repository.GetQuestions()
	if err != nil {
		return nil, err
	}
	return questions, nil
}
