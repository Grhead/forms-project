package service

import (
	"context"
	"time"
	"tusur-forms/internal/domain"
)

type FormUniqResp struct {
	ID            string
	ExternalID    string
	Title         string
	DocumentTitle string
	CreatedAt     time.Time
	Questions     []*domain.Question
	Responses     []*FormResponseUniqResp
}

type FormResponseUniqResp struct {
	ResponseID string
	Answers    map[string]domain.Answer
}

func (f *FormUniqResp) ToDomain() domain.Form {
	// var questions = make([]*domain.Question, 0, len(f.Responses))
	// for _, i := range f.Responses {
	// 	questions = append(questions, i.Questions...)
	// }
	return domain.Form{
		ID:            f.ID,
		ExternalID:    f.ExternalID,
		Title:         f.Title,
		DocumentTitle: f.DocumentTitle,
		CreatedAt:     f.CreatedAt,
		Questions:     f.Questions,
	}
}

// func (q *QuestionUniqResp) ToDomain() domain.Question {
// 	var answers []*domain.Answer
// 	for _, i := range q.Responses {
// 		answers = append(answers, i.Answers...)
// 	}
// 	return domain.Question{
// 		Title:           q.Title,
// 		Description:     q.Description,
// 		Type:            q.Type,
// 		IsRequired:      q.IsRequired,
// 		Answers:         answers,
// 		PossibleAnswers: q.PossibleAnswers,
// 	}
// }

type FormServiceProvider interface {
	NewService(ctx context.Context, filename string) (FormService, error)
}

type FormService interface {
	NewForm(title string, documentTitle string) (domain.Form, error)
	GetForm(formID string) (*FormUniqResp, error)
	SetQuestions(form domain.Form, questions []*domain.Question) (*domain.Form, error)
	GetResponseList(externalID string, questionID string, keyQID string) ([]*FormResponseUniqResp, error)
}
