package database

import "time"

type dbAnswer struct {
	AnswerId        string `gorm:"primaryKey"`
	AnswerTimestamp time.Time
	AnswerContent   string
	FormsQuestionId string
	FormsQuestion   dbFormsQuestion `gorm:"references:FormsQuestionId"`
}
