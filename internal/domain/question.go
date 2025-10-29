package domain

type Question struct {
	Id              string
	Title           string
	Type            QuestionType
	IsRequired      bool
	PossibleAnswers []PossibleAnswer
}

type QuestionType struct {
	Id    string
	Title string
}
