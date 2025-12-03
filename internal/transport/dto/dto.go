package dto

import (
	"time"
)

type Form struct {
	Title         string      `json:"title"`
	DocumentTitle string      `json:"documentTitle"`
	CreatedAt     time.Time   `json:"createdAt"`
	Questions     []*Question `json:"questions"`
}

type Question struct {
	Title           string            `json:"title"`
	Description     string            `json:"description"`
	Type            string            `json:"type"`
	IsRequired      bool              `json:"isRequired"`
	PossibleAnswers []*PossibleAnswer `json:"possibleAnswers"`
}

type PossibleAnswer struct {
	Content string `json:"content"`
}

type Answer struct {
	SubmittedAt time.Time `json:"submittedAt"`
	Content     string    `json:"content"`
}
