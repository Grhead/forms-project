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

func (f *Form) ToDTO() *dto.ResponseForm {
	if len(f.Questions) != 0 {
		var qs = make([]*dto.ResponseQuestion, 0, len(f.Questions))
		for _, q := range f.Questions {
			qs = append(qs, q.ToDTO())
		}
		return &dto.ResponseForm{
			ExternalID:    f.ExternalID,
			Title:         f.Title,
			DocumentTitle: f.DocumentTitle,
			CreatedAt:     f.CreatedAt,
			Questions:     qs,
		}
	} else {
		return &dto.ResponseForm{
			Title:         f.Title,
			DocumentTitle: f.DocumentTitle,
			CreatedAt:     f.CreatedAt,
		}
	}
}
