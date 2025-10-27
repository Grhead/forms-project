package database

type dbQuestion struct {
	QuestionId               string `gorm:"primaryKey"`
	QuestionTitle            string
	QuestionTypeId           []dbQuestionType `gorm:"foreignKey:QuestionTypeId"`
	IsRequired               bool
	QuestionPossibleAnswerId []dbQuestionPossibleAnswer `gorm:"foreignKey:QuestionPossibleAnswerId"`
}

type dbQuestionType struct {
	QuestionTypeId    string `gorm:"primaryKey"`
	QuestionTypeTitle string
}
