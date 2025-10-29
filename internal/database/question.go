package database

type dbQuestion struct {
	Id               string `gorm:"primaryKey"`
	Title            string
	TypeId           []dbQuestionType `gorm:"foreignKey:Id"`
	IsRequired       bool
	PossibleAnswerId []dbQuestionPossibleAnswer `gorm:"foreignKey:Id"`
}

type dbQuestionType struct {
	Id    string `gorm:"primaryKey"`
	Title string
}
