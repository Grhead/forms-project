package domain

import "time"

type Form struct {
	Id            string
	ExternalId    string
	Title         string
	DocumentTitle string
	CreatedAt     time.Time
	Questions     []*Question
	Answers       []*Answer
}
