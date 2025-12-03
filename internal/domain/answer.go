package domain

import (
	"time"
	"tusur-forms/internal/transport/dto"
)

type Answer struct {
	SubmittedAt time.Time
	Content     string
}

func (a *Answer) ToDTO() *dto.Answer {
	return &dto.Answer{
		SubmittedAt: a.SubmittedAt,
		Content:     a.Content,
	}
}
