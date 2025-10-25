package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/forms/v1"
	"google.golang.org/api/option"
)

func createAuthLink(conf *oauth2.Config) {
	authURL := conf.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)
}

func exchangeToken(ctx context.Context, conf *oauth2.Config, code string) (*oauth2.Token, error) {
	tok, err := conf.Exchange(ctx, code)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok, nil
}

func saveToken(token *oauth2.Token) error {
	f, err := os.Create("token.json")
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatalf("Unable to cache oauth token: %v", err)
		}
	}(f)
	err = json.NewEncoder(f).Encode(token)
	if err != nil {
		return err
	}
	return nil
}

func tokenFromFile(file string) *oauth2.Token {
	f, err := os.Open(file)
	if err != nil {
		log.Fatalf("Unable to open token file: %v", err)
		return nil
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatalf("Unable to cache token file: %v", err)
		}
	}(f)
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok
}
package main

import (
"context"
"encoding/json"
"log"
"os"

"github.com/joho/godotenv"
"golang.org/x/oauth2"
"golang.org/x/oauth2/google"
"google.golang.org/api/forms/v1"
"google.golang.org/api/option"
)

func main() {
	ctx := context.Background()
	var config = &oauth2.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		Endpoint:     google.Endpoint,
		Scopes:       []string{forms.FormsBodyScope, forms.FormsResponsesReadonlyScope},
		RedirectURL:  redirectUrl,
	}

	tok := tokenFromFile("token.json")
	TokenSource := config.TokenSource(ctx, tok)

	svc, err := forms.NewService(ctx, option.WithTokenSource(TokenSource))
	if err != nil {
		log.Fatalf("Unable to retrieve Spaces service %v", err)
	}
	/*var formConfig = &forms.Form{
		Info: &forms.Info{
			Title:         "Форма обратной связи по мероприятию",
			DocumentTitle: "Отзыв о конференции TechCon 2025",
		},
	}*/

	itemsToAdd := []*forms.Item{
		{
			Title: "Как вас зовут?",
			QuestionItem: &forms.QuestionItem{
				Question: &forms.Question{
					QuestionId:   "1001",
					TextQuestion: &forms.TextQuestion{},
				},
			},
		},
		{
			Title: "Какова ваша общая оценка мероприятия?",
			QuestionItem: &forms.QuestionItem{
				Question: &forms.Question{
					QuestionId: "1002",
					ChoiceQuestion: &forms.ChoiceQuestion{
						Type: "RADIO",
						Options: []*forms.Option{
							{Value: "Отлично"},
							{Value: "Хорошо"},
							{Value: "Удовлетворительно"},
							{Value: "Плохо"},
						},
					},
				},
			},
		},
	}
	var requests []*forms.Request
	for i, item := range itemsToAdd {
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
	/*id, err := svc.Forms.Create(formConfig).Do()
	fmt.Println(id.FormId)
	_, err = svc.Forms.BatchUpdate("10zLnhdRl84-poEbECNzTFcpKcYXfnaSoCXoX8vNorG8", &forms.BatchUpdateFormRequest{
		Requests: requests,
	}).Do()
	if err != nil {
		log.Fatalf("Unable to create form: %v", err)
	}*/
	do, err := svc.Forms.Responses.Get("10zLnhdRl84-poEbECNzTFcpKcYXfnaSoCXoX8vNorG8", "ACYDBNiM1N6j4QdrilDTpgVTSkKATHRYAtblFpOQk8vRETDevLlA2_Fii-gSWHEJmJGZYAU").Do()
	if err != nil {
		log.Fatalf("Unable to get form: %v", err)
	}
	for questionID, answer := range do.Answers {
		if answer.TextAnswers != nil && len(answer.TextAnswers.Answers) > 0 {
			answerValue := answer.TextAnswers.Answers[0].Value

			log.Printf("ID Вопроса: %s (Ответ): %s", questionID, answerValue)
		} else {
			log.Printf("ID Вопроса: %s: Ответ не является текстовым или пустым. Пропуск.", questionID)
		}
	}
	log.Println("---------------------------------------")
}

func tokenFromFile(file string) *oauth2.Token {
	f, err := os.Open(file)
	if err != nil {
		log.Fatalf("Unable to open token file: %v", err)
		return nil
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatalf("Unable to close token file: %v", err)
		}
	}(f)
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok
}
