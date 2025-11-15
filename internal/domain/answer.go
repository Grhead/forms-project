package domain

import "time"

type Answer struct {
	ID          string
	SubmittedAt time.Time
	Content     string
	Form      Form
	Question  Question
}

type PossibleAnswer struct {
	Content string
}
