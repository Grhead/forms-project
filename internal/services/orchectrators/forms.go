package orchectrators

import (
	"log"
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
		err = s.creator.SetQuestions(d, questions[0])

		if err != nil {
			return nil, err
		}
		for i := range questions[0] {
			log.Println(questions[0][i])
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
	form, err = s.repository.GetForm(d.ID)
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
	for _, item := range form.Questions {
		for _, f := range item.Answers {
			exists, err := s.repository.CheckResponseEnvironmentExists(f.ResponseID)
			if err != nil {
				return nil, err
			}
			if exists {
				continue
			}
			log.Println(f.ResponseID)

			err = s.repository.CreateAnswer(f.ToDomain(), formID, item.ID, f.ResponseID)
			if err != nil {
				return nil, err
			}
		}
		//for key, answer := range item.Answers {
		//	log.Println("KEY " + key + "ANSWER " + answer.Content)
		//}
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
