package google

import (
	"context"
	"log"
	"time"
	"tusur-forms/internal/database"
	"tusur-forms/internal/domain"
	service "tusur-forms/internal/services/forms"

	"github.com/google/uuid"
	"golang.org/x/oauth2"
	"google.golang.org/api/forms/v1"
	"google.golang.org/api/option"
)

type GoogleForms struct {
	TokenSource oauth2.TokenSource
}
type googleFormsAdapter struct {
	googleClient *forms.Service
	repository   database.FormRepository
}

func (g *GoogleForms) NewService(ctx context.Context, r database.FormRepository) (service.FormService, error) {
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

func (g *googleFormsAdapter) GetForm(formID string) (*service.FormUniqResp, error) {
	externalID, err := g.repository.GetFormExternalID(formID)
	if err != nil {
		return nil, err
	}
	resultForm, err := g.googleClient.Forms.Get(externalID).Do()
	if err != nil {
		return nil, err
	}
	responseList, err := g.googleClient.Forms.Responses.List(externalID).Do()
	if err != nil {
		return nil, err
	}
	var questions []*service.QuestionUniqResp

	for _, item := range resultForm.Items {
		if item.QuestionItem == nil {
			continue
		}
		googleQID := item.QuestionItem.Question.QuestionId
		qInternalID, err := g.repository.GetQuestionIDByTitle(item.Title)
		if err != nil || qInternalID == "" {
			return nil, err
		}
		tempQuestion := service.QuestionUniqResp{
			ID:          qInternalID,
			Title:       item.Title,
			Description: item.Description,
			Type:        domain.QuestionType{},
			IsRequired:  item.QuestionItem.Question.Required,
			Answers:     make([]*service.AnswerUniq, 0),
		}
		if item.QuestionItem.Question.ChoiceQuestion != nil {
			for _, opt := range item.QuestionItem.Question.ChoiceQuestion.Options {
				tempQuestion.PossibleAnswers = append(tempQuestion.PossibleAnswers, &domain.PossibleAnswer{
					Content: opt.Value,
				})
			}
		}
		for _, resp := range responseList.Responses {
			if answerItem, ok := resp.Answers[googleQID]; ok {
				if answerItem.TextAnswers != nil && len(answerItem.TextAnswers.Answers) > 0 {
					for _, textAnswer := range answerItem.TextAnswers.Answers {
						t, _ := time.Parse(time.RFC3339, resp.CreateTime)
						domainAns := service.AnswerUniq{
							ResponseID:  resp.ResponseId,
							Content:     textAnswer.Value,
							SubmittedAt: t,
						}
						tempQuestion.Answers = append(tempQuestion.Answers, &domainAns)
					}
				} else {
					t, _ := time.Parse(time.RFC3339, resp.CreateTime)
					domainAns := service.AnswerUniq{
						ResponseID:  resp.ResponseId,
						Content:     "",
						SubmittedAt: t,
					}
					tempQuestion.Answers = append(tempQuestion.Answers, &domainAns)

				}
			}
		}

		questions = append(questions, &tempQuestion)
	}

	f := service.FormUniqResp{
		ID:            formID,
		ExternalID:    externalID,
		Title:         resultForm.Info.Title,
		DocumentTitle: resultForm.Info.DocumentTitle,
		Questions:     questions,
	}

	return &f, nil
}

func (g *googleFormsAdapter) SetQuestions(form domain.Form, questions []*domain.Question) error {
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
		return err
	}
	log.Println("BOOM W")
	return nil
}

//func (g *googleFormsAdapter) GetResponse(formID string, responseID string, questionID string, keyQID string) (service.FormResponseUniqResp, error) {
//	AnswerDTO := service.FormResponseUniqResp{
//		ResponseID: responseID,
//		Answers:    make(map[string]domain.Answer),
//	}
//
//	response, err := g.googleClient.Forms.Responses.Get(formID, responseID).Do()
//	if err != nil {
//		return service.FormResponseUniqResp{}, err
//	}
//	if item, ok := response.Answers[questionID]; ok {
//		submittedTime, _ := time.Parse(time.RFC3339, response.CreateTime)
//
//		if item.TextAnswers == nil || len(item.TextAnswers.Answers) == 0 {
//			AnswerDTO.Answers[keyQID] = domain.Answer{
//				SubmittedAt: submittedTime,
//				Content:     "",
//			}
//		} else {
//			var values []string
//			for _, ta := range item.TextAnswers.Answers {
//				values = append(values, ta.Value)
//			}
//
//			content := strings.Join(values, ", ")
//
//			AnswerDTO.Answers[keyQID] = domain.Answer{
//				SubmittedAt: submittedTime,
//				Content:     content,
//			}
//		}
//	}
//	// for _, item := range responses {
//	// 	for key, answer := range item.Answers {
//	// 		log.Println("KEY " + key + "ANSWER " + answer.Content)
//	// 	}
//	// }
//	return AnswerDTO, nil
//}

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
