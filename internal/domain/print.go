package domain

import (
	"fmt"
)

type IOutput interface {
	Print() string
}

func (f *Form) Print() string {
	var questions []string
	if len(f.Questions) != 0 {
		for _, item := range f.Questions {
			questions = append(questions, item.Print())
		}
	}
	var result = fmt.Sprintf("ID: %s ExternalID: %s Title: %s DocumentTitle: %s CreatedAt: %s Questions: %s",
		f.ID, f.ExternalID, f.Title, f.DocumentTitle, f.CreatedAt, questions)
	return result
}
func (q *Question) Print() string {
	var answers []string
	for _, item := range q.Answers {
		answers = append(answers, item.Print())
	}
	var possibleAnswers []string
	for _, item := range q.PossibleAnswers {
		possibleAnswers = append(possibleAnswers, item.Print())
	}
	var result = fmt.Sprintf("ID: %s Title: %s Description: %s Type: %s IsRequired: %t Answers: %v PossibleAnswers: %v",
		q.ID, q.Title, q.Description, q.Type.Print(), q.IsRequired, answers, possibleAnswers)
	return result
}
func (p *PossibleAnswer) Print() string {
	var result = fmt.Sprintf("Content: %s", p.Content)
	return result
}
func (a *Answer) Print() string {
	var result = fmt.Sprintf("ID: %s SubmittedAt: %s Content: %s",
		a.ID, a.SubmittedAt, a.Content)
	return result
}
func (t *QuestionType) Print() string {
	var result = fmt.Sprintf("ID: %s Title: %s",
		t.ID, t.Title)
	return result
}
