package google

import (
	"context"
	"log/slog"
	"time"
	"tusur-forms/internal/domain"
	"tusur-forms/internal/repository"
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
	repository   repository.FormRepository
}

func (g *GoogleForms) NewService(ctx context.Context, r repository.FormRepository) (service.FormService, error) {
	svc, err := forms.NewService(ctx, option.WithTokenSource(g.TokenSource))
	if err != nil {
		slog.Error("Failed to create new GoogleForms service", "err", err)
		return nil, err
	}
	adapter := &googleFormsAdapter{
		googleClient: svc,
		repository:   r,
	}
	return adapter, nil
}

func (g *googleFormsAdapter) NewForm(title string, documentTitle string, description string) (domain.Form, error) {
	gf := &forms.Form{
		Info: &forms.Info{
			Title:         title,
			DocumentTitle: documentTitle,
		},
	}
	result, err := g.googleClient.Forms.Create(gf).Do()
	if err != nil {
		slog.Error("Failed to create new GoogleForm", "err", err)
		return domain.Form{}, err
	}
	var requests = make([]*forms.Request, 0, 1)

	requests = append(requests, &forms.Request{
		UpdateFormInfo: &forms.UpdateFormInfoRequest{
			Info: &forms.Info{
				Description: description,
			},
			UpdateMask: "Description",
		},
	})
	_, err = g.googleClient.Forms.BatchUpdate(
		result.FormId,
		&forms.BatchUpdateFormRequest{Requests: requests}).
		Do()
	if err != nil {
		slog.Error("Failed to batch update form", "external_id", result.FormId, "err", err)
		return domain.Form{}, err
	}
	f := domain.Form{
		ID:            uuid.NewString(),
		ExternalID:    result.FormId,
		Title:         result.Info.Title,
		DocumentTitle: result.Info.DocumentTitle,
		Description:   result.Info.Description,
		CreatedAt:     time.Now(),
		Questions:     nil,
	}
	return f, nil
}

func (g *googleFormsAdapter) GetForm(formID string) (*service.FormUniqResp, error) {
	externalID, err := g.repository.GetFormExternalID(formID)
	if err != nil {
		slog.Error("Failed to get external id from repository", "form_id", formID, "error", err)
		return nil, err
	}
	resultForm, err := g.googleClient.Forms.Get(externalID).Do()
	if err != nil {
		slog.Error("Failed to get form", "external_id", externalID, "error", err)
		return nil, err
	}
	responseList, err := g.googleClient.Forms.Responses.List(externalID).Do()
	if err != nil {
		slog.Error("Failed to get responses list", "external_id", externalID, "error", err)
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
			slog.Error("Failed to get question id by title", "question_title", item.Title, "form_id", formID, "error", err)
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

func (g *googleFormsAdapter) SetQuestions(form *domain.Form, questions []*domain.Question) error {
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
		slog.Error("Failed to batch update form", "external_id", form.ExternalID, "internal_id", form.ID, "err", err)
		return err
	}
	return nil
}
