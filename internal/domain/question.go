package domain

type Question struct {
	Id              string
	Title           string
	Description     string
	Type            QuestionType
	IsRequired      bool
	PossibleAnswers []PossibleAnswer
}

type QuestionType struct {
	Id    string
	Title QuestionTypeTitles
}

type QuestionTypeTitles string

const TypeCheckbox QuestionTypeTitles = "CHECKBOX"
const TypeRadio QuestionTypeTitles = "RADIO"

const TypeText QuestionTypeTitles = "TEXT"

//const TypeDate QuestionTypeTitles = "DATE"
//const ScaleQuestion QuestionTypeTitles = "SCALE"
