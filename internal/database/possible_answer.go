package database

type dbPossibleAnswer struct {
	PossibleAnswerId      string `gorm:"primaryKey"`
	PossibleAnswerContent string
}

type dbQuestionPossibleAnswer struct {
	QuestionPossibleAnswerId string `gorm:"primaryKey"`
	QuestionId               string
	Question                 dbQuestion `gorm:"references:QuestionId"`
	PossibleAnswerId         string
	PossibleAnswer           dbPossibleAnswer `gorm:"references:PossibleAnswerId"`
}
