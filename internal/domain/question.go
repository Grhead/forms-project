package domain

type Question struct {
	ID              string
	Title           string
	Description     string
	Type            QuestionType
	IsRequired      bool
	Answers         []*Answer
	PossibleAnswers []*PossibleAnswer
}

type QuestionType struct {
	ID    string
	Title QuestionTypeTitles
}

type QuestionTypeTitles string

const TypeCheckbox QuestionTypeTitles = "CHECKBOX"
const TypeRadio QuestionTypeTitles = "RADIO"
const TypeText QuestionTypeTitles = "TEXT"

//const TypeDate QuestionTypeTitles = "DATE"
//const ScaleQuestion QuestionTypeTitles = "SCALE"
