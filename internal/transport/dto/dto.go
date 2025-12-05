package dto

import (
	"time"
)

type ResponseForm struct {
	ExternalID    string              `json:"externalID"`
	Title         string              `json:"title"`
	DocumentTitle string              `json:"documentTitle"`
	Description   string              `json:"description"`
	CreatedAt     time.Time           `json:"createdAt"`
	Questions     []*ResponseQuestion `json:"questions"`
}

type ResponseQuestion struct {
	Title           string            `json:"title"`
	Description     string            `json:"description"`
	Type            string            `json:"type"`
	IsRequired      bool              `json:"isRequired"`
	PossibleAnswers []*PossibleAnswer `json:"possibleAnswers"`
}

type RequestForm struct {
	Title         string   `json:"title"`
	DocumentTitle string   `json:"documentTitle"`
	Description   string   `json:"description"`
	Questions     []string `json:"questions"`
}

type RequestQuestion struct {
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
