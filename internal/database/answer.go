package database

import (
	"time"
	"tusur-forms/internal/domain"

	"gorm.io/gorm"
)

type dbAnswer struct {
	Id              string `gorm:"primaryKey"`
	SubmittedAt     time.Time
	Content         string
	FormsQuestionId string
	FormsQuestion   dbFormsQuestion `gorm:"foreignKey:FormsQuestionId;references:Id"`
}

func CreateAnswer(a *domain.Answer, db *gorm.DB) error {
	fq, err := getFormsQuestionId(a, db)
	if err != nil {
		return err
	}
	dbQ := dbAnswer{
		Id:              a.Id,
		SubmittedAt:     a.SubmittedAt,
		Content:         a.Content,
		FormsQuestionId: fq.Id,
		FormsQuestion:   dbFormsQuestion{},
	}
	err = db.Create(&dbQ).Error
	if err != nil {
		return err
	}
	return db.Save(&dbQ).Error
}

func getFormsQuestionId(a *domain.Answer, db *gorm.DB) (dbFormsQuestion, error) {
	var fq dbFormsQuestion
	err := db.Where("form_id = ? AND question_id = ?", a.FormId, a.QuestionId).First(&fq).Error
	return fq, err

}
