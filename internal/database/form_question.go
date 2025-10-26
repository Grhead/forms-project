package database

type dbFormQuestion struct {
	FormsQuestionId string
	FormId          []dbForm
	QuestionId      []dbQuestion
}
