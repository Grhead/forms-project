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
	Questions     []*QuestionUniqResp
}

type QuestionUniqResp struct {
	ID              string
	Title           string
	Description     string
	Type            domain.QuestionType
	IsRequired      bool
	Answers         []*AnswerUniq
	PossibleAnswers []*domain.PossibleAnswer
}

type AnswerUniq struct {
	ResponseID  string
	SubmittedAt time.Time
	Content     string
}

func (f *FormUniqResp) ToDomain() *domain.Form {
	var questions = make([]*domain.Question, 0, len(f.Questions))
	for _, i := range f.Questions {
		questions = append(questions, i.ToDomain())
	}
	return &domain.Form{
		ID:            f.ID,
		ExternalID:    f.ExternalID,
		Title:         f.Title,
		DocumentTitle: f.DocumentTitle,
		CreatedAt:     f.CreatedAt,
		Questions:     questions,
	}
}

func (q *QuestionUniqResp) ToDomain() *domain.Question {
	var answers = make([]*domain.Answer, 0, len(q.Answers))
	for _, a := range q.Answers {
		answers = append(answers, a.ToDomain())
	}
	return &domain.Question{
		Title:           q.Title,
		Description:     q.Description,
		Type:            q.Type,
		IsRequired:      q.IsRequired,
		Answers:         answers,
		PossibleAnswers: q.PossibleAnswers,
	}
}

func (q *AnswerUniq) ToDomain() *domain.Answer {
	return &domain.Answer{
		SubmittedAt: q.SubmittedAt,
		Content:     q.Content,
	}
}

type FormServiceProvider interface {
	NewService(ctx context.Context, filename string) (FormService, error)
}

type FormService interface {
	NewForm(title string, documentTitle string) (domain.Form, error)
	GetForm(formID string) (*FormUniqResp, error)
	SetQuestions(form domain.Form, questions []*domain.Question) error
	//GetResponse(formID string, responseID string, questionID string, keyQID string) (FormResponseUniqResp, error)
}
