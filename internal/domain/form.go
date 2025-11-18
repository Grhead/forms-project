package domain

import "time"

type Form struct {
	ID            string
	ExternalID    string
	Title         string
	DocumentTitle string
	CreatedAt     time.Time
	Questions     []*Question
}

