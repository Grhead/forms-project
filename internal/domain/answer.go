package domain

import "time"

type Answer struct {
	SubmittedAt time.Time
	Content     string
}

type PossibleAnswer struct {
	Content string
}

