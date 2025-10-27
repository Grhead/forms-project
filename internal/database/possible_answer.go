package database

type dbPossibleAnswer struct {
	PossibleAnswerId      string `gorm:"primaryKey"`
	PossibleAnswerContent string
}

type dbQuestionPossibleAnswer struct {
	QuestionPossibleAnswerId string             `gorm:"primaryKey"`
	QuestionId               []dbQuestion       `gorm:"foreignKey:QuestionId"`
	PossibleAnswerId         []dbPossibleAnswer `gorm:"foreignKey:PossibleAnswerId"`
}
