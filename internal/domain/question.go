package domain

import "tusur-forms/internal/transport/dto"

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

type PossibleAnswer struct {
	Content string
}

type QuestionTypeTitles string

const TypeCheckbox QuestionTypeTitles = "CHECKBOX"
const TypeRadio QuestionTypeTitles = "RADIO"
const TypeText QuestionTypeTitles = "TEXT"

//const TypeDate QuestionTypeTitles = "DATE"
//const ScaleQuestion QuestionTypeTitles = "SCALE"

func (q *Question) ToDTO() *dto.Question {
	var pa = make([]*dto.PossibleAnswer, 0, len(q.PossibleAnswers))
	for _, p := range q.PossibleAnswers {
		pa = append(pa, &dto.PossibleAnswer{
			Content: p.Content,
		})
	}
	return &dto.Question{
		Title:           q.Title,
		Description:     q.Description,
		Type:            string(q.Type.Title),
		IsRequired:      q.IsRequired,
		PossibleAnswers: pa,
	}
}

func (p *PossibleAnswer) ToDTO() *dto.PossibleAnswer {
	return &dto.PossibleAnswer{
		Content: p.Content,
	}
}
