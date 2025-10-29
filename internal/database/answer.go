package database

import "time"

type dbAnswer struct {
	Id              string `gorm:"primaryKey"`
	SubmittedAt     time.Time
	Content         string
	FormsQuestionId []dbFormsQuestion `gorm:"foreignKey:Id"`
}
