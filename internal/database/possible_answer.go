package database

type dbPossibleAnswer struct {
	PossibleAnswerId      string
	PossibleAnswerContent string
}

type dbQuestionPossibleAnswer struct {
	QuestionPossibleAnswerId string
	QuestionId               []dbQuestion
	PossibleAnswerId         []dbPossibleAnswer
}
