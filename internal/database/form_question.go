package database

import (
	"github.com/google/uuid"
)

type dbFormsQuestion struct {
	ID         string `gorm:"primaryKey"`
	FormID     string
	Form       dbForm `gorm:"foreignKey:FormID;references:ID"`
	QuestionID string
	Question   dbQuestion `gorm:"foreignKey:QuestionID;references:ID"`
}

func (g *GormRepository) CreateFormsQuestion(fID string, qID string) error {
	dbQ := dbFormsQuestion{
		ID:         uuid.NewString(),
		FormID:     fID,
		Form:       dbForm{},
		QuestionID: qID,
		Question:   dbQuestion{},
	}

	err := g.db.Create(&dbQ).Error
	if err != nil {
		return err
	}
	return nil
}
