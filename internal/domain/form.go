package domain

import (
	"time"
	"tusur-forms/internal/transport/dto"
)

type Form struct {
	ID            string
	ExternalID    string
	Title         string
	DocumentTitle string
	CreatedAt     time.Time
	Questions     []*Question
}

func (f *Form) ToDTO() *dto.Form {
	var qs = make([]*dto.Question, 0, len(f.Questions))
	for _, q := range f.Questions {
		qs = append(qs, q.ToDTO())
	}
	return &dto.Form{
		Title:         f.Title,
		DocumentTitle: f.DocumentTitle,
		CreatedAt:     f.CreatedAt,
		Questions:     qs,
	}
}
