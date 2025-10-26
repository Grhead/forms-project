package database

import "time"

type dbAnswer struct {
	AnswerId        string
	AnswerTimestamp time.Time
	AnswerContent   string
	FormsQuestionId []dbQuestionType
}
