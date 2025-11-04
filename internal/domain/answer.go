package domain

import "time"

type Answer struct {
	Id          string
	SubmittedAt time.Time
	Content     string
	FormId      string
	QuestionId  string
}

type PossibleAnswer struct {
	Content string
}
