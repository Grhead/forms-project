package database

type dbQuestion struct {
	QuestionId       string
	QuestionTitle    string
	QuestionTypeId   []dbQuestionType
	IsRequired       bool
	PossibleAnswerId []dbPossibleAnswer
}

type dbQuestionType struct {
	QuestionTypeId    string
	QuestionTypeTitle string
}
