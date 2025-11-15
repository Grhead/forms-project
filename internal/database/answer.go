package database

import (
	"time"
	"tusur-forms/internal/domain"
)

type dbAnswer struct {
	ID              string `gorm:"primaryKey"`
	SubmittedAt     time.Time
	Content         string
	FormsQuestionID string
	FormsQuestion   dbFormsQuestion `gorm:"foreignKey:FormsQuestionID;references:ID"`
}

func (g *GormRepository) CreateAnswer(a *domain.Answer) error {
	fq, err := g.getFormsQuestionID(a)
	if err != nil {
		return err
	}
	dbQ := dbAnswer{
		ID:              a.ID,
		SubmittedAt:     a.SubmittedAt,
		Content:         a.Content,
		FormsQuestionID: fq.ID,
		FormsQuestion:   dbFormsQuestion{},
	}
	err = g.db.Create(&dbQ).Error
	if err != nil {
		return err
	}
	return g.db.Save(&dbQ).Error
}

func (g *GormRepository) getFormsQuestionID(a *domain.Answer) (dbFormsQuestion, error) {
	// var fq dbFormsQuestion
	// err := g.db.Where("form_id = ? AND question_id = ?", a.FormID, a.QuestionID).First(&fq).Error
	// return fq, err
	return dbFormsQuestion{
		ID:         "",
		FormID:     "",
		Form:       dbForm{},
		QuestionID: "",
		Question:   dbQuestion{},
	}, nil
}
