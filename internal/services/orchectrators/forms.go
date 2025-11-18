package orchectrators

import (
	"tusur-forms/internal/database"
	"tusur-forms/internal/domain"
	service "tusur-forms/internal/services/forms"
)

type FormsOrchestrator struct {
	creator    service.FormService
	repository database.FormRepository
}

func NewFormsOrchestrator(c service.FormService, r database.FormRepository) *FormsOrchestrator {
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
		form, err = s.creator.SetQuestions(d, questions[0])
		if err != nil {
			return nil, err
		}
		for i := range questions[0] {
			qID, err := s.repository.CreateQuestion(questions[0][i])
			if err != nil {
				return nil, err
			}
			err = s.repository.CreateFormsQuestion(form, qID)
			if err != nil {
				return nil, err
			}
		}
		return form, nil
	}
	return form, nil
}

func (s *FormsOrchestrator) CheckoutAnswers(formID string) (*domain.Form, error) {
	form, err := s.creator.GetForm(formID)
	if err != nil {
		return nil, err
	}
	// questions, err := s.repository.GetQuestionIDs(formID)
	// if err != nil {
	// 	return nil, err
	// }
	// var allowResp []string
	for _, item := range form.Responses {
		// exists, err := s.repository.CheckResponseEnvironmentExists(item.ResponseID)
		// if err != nil {
		// 	return nil, err
		// }
		// if exists {
		// 	continue
		// }
				// log.Println(item.ResponseID)

		for key, f := range item.Answers {
				// log.Println(key)
				// log.Println(f)
				s.repository.CreateAnswer(&f, formID, key, item.ResponseID)
			}
		// for key, answer := range item.Answers {
		// 	log.Println("KEY " + key + "ANSWER " + answer.Content)
		// }
	}
	
	// for _, i := range allowResp {
	// 	for qID, question := range form.Questions {
	// 		log.Println(question.Title)
	// 		for _, f := range question.Answers {
	// 			log.Println(f)
	// 			s.repository.CreateAnswer(f, formID, questions[qID], i)
	// 		}
	// 	}
	// }
	domainForm, err := s.repository.GetForm(formID)
	if err != nil {
		return nil, err
	}
	return domainForm, nil
}
