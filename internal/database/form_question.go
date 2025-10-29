package database

type dbFormsQuestion struct {
	FormsQuestionId string       `gorm:"primaryKey"`
	FormId          []dbForm     `gorm:"foreignKey:FormId"`
	QuestionId      []dbQuestion `gorm:"foreignKey:QuestionId"`
}
