package google

import (
	"context"
	"log"
	"time"
	"tusur-forms/internal/database"
	"tusur-forms/internal/domain"

	"github.com/google/uuid"
	"golang.org/x/oauth2"
	"google.golang.org/api/forms/v1"
	"google.golang.org/api/option"
)

type FormServiceProvider interface {
	NewService(ctx context.Context, filename string) (FormService, error)
}

type FormService interface {
	NewForm(title string, documentTitle string) (domain.Form, error)
	GetForm(formExternalID string) (domain.Form, error)
	SetQuestions(form domain.Form, questions []*domain.Question) (domain.Form, error)
	GetResponseList()
	GetResponse()
}

type GoogleForms struct {
	TokenSource oauth2.TokenSource
}
type googleFormsAdapter struct {
	googleClient *forms.Service
	repository   database.FormRepository
}

func (g *GoogleForms) NewService(ctx context.Context, r database.FormRepository) (FormService, error) {
	svc, err := forms.NewService(ctx, option.WithTokenSource(g.TokenSource))
	if err != nil {
		return nil, err
	}
	adapter := &googleFormsAdapter{
		googleClient: svc,
		repository:   r,
	}

	return adapter, nil
}

func (g *googleFormsAdapter) NewForm(title string, documentTitle string) (domain.Form, error) {
	gf := &forms.Form{
		Info: &forms.Info{
			Title:         title,
			DocumentTitle: documentTitle,
		},
	}
	result, err := g.googleClient.Forms.Create(gf).Do()
	if err != nil {
		return domain.Form{}, err
	}

	f := domain.Form{
		ID:            uuid.NewString(),
		ExternalID:    result.FormId,
		Title:         result.Info.Title,
		DocumentTitle: result.Info.DocumentTitle,
		CreatedAt:     time.Now(),
		Questions:     nil,
	}
	return f, nil
}

func (g *googleFormsAdapter) GetForm(formID string) (domain.Form, error) {
	log.Println("GetForm berfore external")
	externalID, err := g.repository.GetFormExternalID(formID)
	if err != nil {
		log.Println("externalID")
		return domain.Form{}, err
	}
	response, err := g.googleClient.Forms.Get(externalID).Do()
	if err != nil {
		log.Println("response")
		return domain.Form{}, err
	}
	questions := make([]*domain.Question, 0, len(response.Items))
	for _, i := range response.Items {
		tempQuestion := domain.Question{
			Title:           i.Title,
			Description:     i.Description,
			Type:            domain.QuestionType{},
			IsRequired:      i.QuestionItem.Question.Required,
			PossibleAnswers: nil,
		}
		if i.QuestionItem.Question.ChoiceQuestion != nil {
			answers := make([]*domain.PossibleAnswer, 0, len(response.Items))

			for _, q := range i.QuestionItem.Question.ChoiceQuestion.Options {
				pAnswer := domain.PossibleAnswer{
					Content: q.Value,
				}
				answers = append(answers, &pAnswer)
			}
			tempQuestion.PossibleAnswers = answers
		}
		questions = append(questions, &tempQuestion)

	}
	f := domain.Form{
		ID:            formID,
		ExternalID:    externalID,
		Title:         response.Info.Title,
		DocumentTitle: response.Info.DocumentTitle,
		CreatedAt:     time.Time{},
		Questions:     questions,
	}
	return f, nil
}
func (g *googleFormsAdapter) SetQuestions(form domain.Form, questions []*domain.Question) (domain.Form, error) {
	var formItems []*forms.Item
	var requests = make([]*forms.Request, 0, len(formItems))
	for _, question := range questions {
		tempItem := &forms.Item{
			Description:  question.Description,
			Title:        question.Title,
			QuestionItem: &forms.QuestionItem{},
		}
		googleQuestion := forms.Question{
			Required: question.IsRequired,
		}

		switch question.Type.Title {
		case domain.TypeText:
			googleQuestion.TextQuestion = &forms.TextQuestion{}
		case domain.TypeCheckbox:
			var opts = make([]*forms.Option, 0, len(question.PossibleAnswers))
			for _, pa := range question.PossibleAnswers {
				opts = append(opts, &forms.Option{Value: pa.Content})
			}
			googleQuestion.ChoiceQuestion = &forms.ChoiceQuestion{
				Type:    string(domain.TypeCheckbox),
				Options: opts,
			}
		case domain.TypeRadio:
			var opts = make([]*forms.Option, 0, len(question.PossibleAnswers))
			for _, pa := range question.PossibleAnswers {
				opts = append(opts, &forms.Option{Value: pa.Content})
			}
			googleQuestion.ChoiceQuestion = &forms.ChoiceQuestion{
				Type:    string(domain.TypeRadio),
				Options: opts,
			}

		}
		tempItem.QuestionItem.Question = &googleQuestion
		formItems = append(formItems, tempItem)
	}
	for i, item := range formItems {
		requests = append(requests, &forms.Request{
			CreateItem: &forms.CreateItemRequest{
				Item: item,
				Location: &forms.Location{
					Index:           int64(i),
					ForceSendFields: []string{"Index"},
				},
			},
		})
	}
	_, err := g.googleClient.Forms.BatchUpdate(
		form.ExternalID,
		&forms.BatchUpdateFormRequest{Requests: requests}).
		Do()
	if err != nil {
		return domain.Form{}, err
	}
	result, err := g.GetForm(form.ID)
	if err != nil {
		return domain.Form{}, err
	}
	return result, nil
}
func (g *googleFormsAdapter) GetResponse()     {} //TODO create test
func (g *googleFormsAdapter) GetResponseList() {} //TODO create test

//do, err := svc.Forms.Responses.Get("10zLnhdRl84-poEbECNzTFcpKcYXfnaSoCXoX8vNorG8", "ACYDBNiM1N6j4QdrilDTpgVTSkKATHRYAtblFpOQk8vRETDevLlA2_Fii-gSWHEJmJGZYAU").Do()
//if err != nil {
//log.Fatalf("Unable to get form: %v", err)
//}
//for questionID, answer := range do.Answers {
//if answer.TextAnswers != nil && len(answer.TextAnswers.Answers) > 0 {
//answerValue := answer.TextAnswers.Answers[0].Value
//
//log.Printf("ID Вопроса: %s (Ответ): %s", questionID, answerValue)
//} else {
//log.Printf("ID Вопроса: %s: Ответ не является текстовым или пустым. Пропуск.", questionID)
//}
//}
//log.Println("---------------------------------------")
//}
