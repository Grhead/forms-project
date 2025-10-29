package database

type dbFormsQuestion struct {
	FormsQuestionId string `gorm:"primaryKey"`
	FormId          string
	Form            dbForm `gorm:"references:FormId"`
	QuestionId      string
	Question        dbQuestion `gorm:"references:QuestionId"`
}
