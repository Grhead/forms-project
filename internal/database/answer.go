package database

import (
	"time"
	"tusur-forms/internal/domain"

	"gorm.io/gorm"
)

type dbAnswer struct {
	ID              string `gorm:"primaryKey"`
	SubmittedAt     time.Time
	Content         string
	FormsQuestionID string
	FormsQuestion   dbFormsQuestion `gorm:"foreignKey:FormsQuestionID;references:ID"`
}

func CreateAnswer(a *domain.Answer, db *gorm.DB) error {
	fq, err := getFormsQuestionID(a, db)
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
	err = db.Create(&dbQ).Error
	if err != nil {
		return err
	}
	return db.Save(&dbQ).Error
}

func getFormsQuestionID(a *domain.Answer, db *gorm.DB) (dbFormsQuestion, error) {
	var fq dbFormsQuestion
	err := db.Where("form_id = ? AND question_id = ?", a.FormID, a.QuestionID).First(&fq).Error
	return fq, err

}
