package database

import "time"

type dbAnswer struct {
	AnswerId        string `gorm:"primaryKey"`
	AnswerTimestamp time.Time
	AnswerContent   string
	FormsQuestionId []dbFormsQuestion `gorm:"foreignKey:FormsQuestionId"`
}
