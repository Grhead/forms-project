package database

type dbPossibleAnswer struct {
	Id      string `gorm:"primaryKey"`
	Content string
}

type dbQuestionPossibleAnswer struct {
	Id               string             `gorm:"primaryKey"`
	QuestionId       []dbQuestion       `gorm:"foreignKey:Id"`
	PossibleAnswerId []dbPossibleAnswer `gorm:"foreignKey:Id"`
}
