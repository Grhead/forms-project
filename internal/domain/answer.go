package domain

import "time"

type Answer struct {
	ID          string
	SubmittedAt time.Time
	Content     string
}

type PossibleAnswer struct {
	Content string
}
