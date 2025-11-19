package domain

type Question struct {
	Title           string
	Description     string
	Type            QuestionType
	IsRequired      bool
	Answers         []*Answer
	PossibleAnswers []*PossibleAnswer
}

type QuestionType struct {
	Title QuestionTypeTitles
}

type QuestionTypeTitles string

const TypeCheckbox QuestionTypeTitles = "CHECKBOX"
const TypeRadio QuestionTypeTitles = "RADIO"
const TypeText QuestionTypeTitles = "TEXT"

//const TypeDate QuestionTypeTitles = "DATE"
//const ScaleQuestion QuestionTypeTitles = "SCALE"
