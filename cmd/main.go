// package main

// import (
// 	"context"
// 	"encoding/json"
// 	"log"
// 	"os"

// 	"github.com/joho/godotenv"
// 	"golang.org/x/oauth2"
// 	"golang.org/x/oauth2/google"
// 	"google.golang.org/api/forms/v1"
// 	"google.golang.org/api/option"
// )

// func main() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}

// 	clientId := os.Getenv("CLIENT_ID")
// 	clientSecret := os.Getenv("CLIENT_SECRET")
// 	redirectUrl := os.Getenv("REDIRECT_URL")
// 	ctx := context.Background()
// 	var config = &oauth2.Config{
// 		ClientID:     clientId,
// 		ClientSecret: clientSecret,
// 		Endpoint:     google.Endpoint,
// 		Scopes:       []string{forms.FormsBodyScope, forms.FormsResponsesReadonlyScope},
// 		RedirectURL:  redirectUrl,
// 	}

// 	tok := tokenFromFile("token.json")
// 	TokenSource := config.TokenSource(ctx, tok)

// 	svc, err := forms.NewService(ctx, option.WithTokenSource(TokenSource))
// 	if err != nil {
// 		log.Fatalf("Unable to retrieve Spaces service %v", err)
// 	}
// 	/*var formConfig = &forms.Form{
// 		Info: &forms.Info{
// 			Title:         "Форма обратной связи по мероприятию",
// 			DocumentTitle: "Отзыв о конференции TechCon 2025",
// 		},
// 	}*/

// 	itemsToAdd := []*forms.Item{
// 		{
// 			Title: "Как вас зовут?",
// 			QuestionItem: &forms.QuestionItem{
// 				Question: &forms.Question{
// 					QuestionId:   "1001",
// 					TextQuestion: &forms.TextQuestion{},
// 				},
// 			},
// 		},
// 		{
// 			Title: "Какова ваша общая оценка мероприятия?",
// 			QuestionItem: &forms.QuestionItem{
// 				Question: &forms.Question{
// 					QuestionId: "1002",
// 					ChoiceQuestion: &forms.ChoiceQuestion{
// 						Type: "RADIO",
// 						Options: []*forms.Option{
// 							{Value: "Отлично"},
// 							{Value: "Хорошо"},
// 							{Value: "Удовлетворительно"},
// 							{Value: "Плохо"},
// 						},
// 					},
// 				},
// 			},
// 		},
// 	}
// 	var requests []*forms.Request
// 	for i, item := range itemsToAdd {
// 		requests = append(requests, &forms.Request{
// 			CreateItem: &forms.CreateItemRequest{
// 				Item: item,
// 				Location: &forms.Location{
// 					Index:           int64(i),
// 					ForceSendFields: []string{"Index"},
// 				},
// 			},
// 		})
// 	}
// 	/*id, err := svc.Forms.Create(formConfig).Do()
// 	fmt.Println(id.FormId)
// 	_, err = svc.Forms.BatchUpdate("10zLnhdRl84-poEbECNzTFcpKcYXfnaSoCXoX8vNorG8", &forms.BatchUpdateFormRequest{
// 		Requests: requests,
// 	}).Do()
// 	if err != nil {
// 		log.Fatalf("Unable to create form: %v", err)
// 	}*/
// 	do, err := svc.Forms.Responses.Get("10zLnhdRl84-poEbECNzTFcpKcYXfnaSoCXoX8vNorG8", "ACYDBNiM1N6j4QdrilDTpgVTSkKATHRYAtblFpOQk8vRETDevLlA2_Fii-gSWHEJmJGZYAU").Do()
// 	if err != nil {
// 		log.Fatalf("Unable to get form: %v", err)
// 	}
// 	for questionID, answer := range do.Answers {
// 		if answer.TextAnswers != nil && len(answer.TextAnswers.Answers) > 0 {
// 			answerValue := answer.TextAnswers.Answers[0].Value

// 			log.Printf("ID Вопроса: %s (Ответ): %s", questionID, answerValue)
// 		} else {
// 			log.Printf("ID Вопроса: %s: Ответ не является текстовым или пустым. Пропуск.", questionID)
// 		}
// 	}
// 	log.Println("---------------------------------------")
// }

//	func tokenFromFile(file string) *oauth2.Token {
//		f, err := os.Open(file)
//		if err != nil {
//			log.Fatalf("Unable to open token file: %v", err)
//			return nil
//		}
//		defer func(f *os.File) {
//			err := f.Close()
//			if err != nil {
//				log.Fatalf("Unable to close token file: %v", err)
//			}
//		}(f)
//		tok := &oauth2.Token{}
//		err = json.NewDecoder(f).Decode(tok)
//		return tok
//	}
package main

import (
	"fmt"

	textrank "github.com/DavidBelicza/TextRank/v2"
)

func main() {
	rawText := `Посетил 86-100% занятий и выполняю все задания преподавателя.
				Посетил 86-100% занятий и выполняю все задания преподавателя.
				Посетил 86-100% занятий и выполняю все задания преподавателя.
				Пропустил или не выполнил 15-49% занятий / заданий.
				Посетил 86-100% занятий и выполняю все задания преподавателя.
				Посетил 86-100% занятий и выполняю все задания преподавателя.
				Посетил 86-100% занятий и выполняю все задания преподавателя.
				Посетил 86-100% занятий и выполняю все задания преподавателя.
				Посетил 86-100% занятий и выполняю все задания преподавателя.
				Посетил 86-100% занятий и выполняю все задания преподавателя.
				Посетил 86-100% занятий и выполняю все задания преподавателя.
				Посетил 86-100% занятий и выполняю все задания преподавателя.
				Посетил 86-100% занятий и выполняю все задания преподавателя.
				Посетил 86-100% занятий и выполняю все задания преподавателя.
				Посетил 86-100% занятий и выполняю все задания преподавателя.
				Посетил 86-100% занятий и выполняю все задания преподавателя.
				Посетил 86-100% занятий и выполняю все задания преподавателя.
				Посетил 86-100% занятий и выполняю все задания преподавателя.
				Посетил 86-100% занятий и выполняю все задания преподавателя.
				Посетил 86-100% занятий и выполняю все задания преподавателя.
				Посетил 86-100% занятий и выполняю все задания преподавателя.
				`
	// TextRank object
	tr := textrank.NewTextRank()
	// Default Rule for parsing.
	rule := textrank.NewDefaultRule()
	// Default Language for filtering stop words.
	language := textrank.NewDefaultLanguage()
	algorithmDef := textrank.NewDefaultAlgorithm()
	// Default algorithm for ranking text.
	// Add Spanish stop words (just some example).
	language.SetWords("ru", []string{
		// Местоимения
		"я", "ты", "он", "она", "оно", "мы", "вы", "они",
		"меня", "тебя", "его", "её", "нас", "вас", "их",
		"мне", "тебе", "ему", "ей", "нам", "вам", "им",
		"мной", "тобой", "им", "ей", "нами", "вами", "ими",
		"моя", "твоя", "его", "её", "наша", "ваша", "их",
		"моё", "твоё", "наше", "ваше",
		"мои", "твои", "его", "её", "наши", "ваши", "их",
		"себя", "себе", "собой",

		// Артикли и указательные местоимения (в русском - указательные)
		"это", "то", "такой", "такая", "такое", "такие",
		"этот", "эта", "это", "эти",
		"тот", "та", "то", "те",
		"вот",

		// Предлоги
		"в", "во", "на", "с", "со", "из", "изо", "от", "ото",
		"до", "по", "под", "подо", "над", "за", "при", "про",
		"к", "ко", "у", "об", "обо", "но", "о", "об", "обо",
		"из-за", "из-под", "по-над", "по-за",

		// Союзы
		"и", "или", "либо", "ни", "но", "да", "однако",
		"зато", "тоже", "также", "причем", "причём",
		"как", "что", "чтобы", "чтоб", "будто", "словно",
		"точно", "если", "ежели", "когда", "пока", "хотя",
		"пусть", "дабы", "ибо", "поскольку", "так как",
		"потому что", "оттого что", "вследствие того что",

		// Частицы
		"ли", "бы", "б", "же", "ж", "ведь", "вот", "дескать",
		"мол", "не", "ни", "ну", "уж", "точно", "просто",
		"прямо", "именно", "почти", "единственно", "только",
		"лишь", "исключительно", "почти", "едва", "чуть",

		// Междометия
		"ах", "ох", "эй", "ого", "увы", "фи", "тьфу", "брр",

		// Вспомогательные глаголы и связки
		"быть", "есть", "был", "была", "было", "были",
		"стать", "стал", "стала", "стало", "стали",
		"являться", "является", "являлся", "являлась",
		"являлось", "являлись",
		"становиться", "становится", "становился",

		// Наречия (частотные)
		"уже", "еще", "ещё", "очень", "слишком", "совсем",
		"вполне", "почти", "точно", "именно", "прямо",
		"сразу", "вдруг", "снова", "опять", "тут", "там",
		"здесь", "везде", "всегда", "иногда", "никогда",
		"очень", "сильно", "много", "мало", "немного",
		"слегка", "чуть", "чуть-чуть",

		// Вопросительные слова
		"кто", "что", "какой", "какая", "какое", "какие",
		"чей", "чья", "чьё", "чьи", "где", "куда", "откуда",
		"когда", "почему", "зачем", "как", "сколько",

		// Отрицания
		"нет", "не", "ни", "никак", "никогда", "никуда",
		"ниоткуда", "нисколько", "ничуть",

		// Числительные (частотные)
		"один", "одна", "одно", "одни", "два", "две",
		"три", "четыре", "пять", "шесть", "семь", "восемь",
		"девять", "десять", "первый", "второй", "третий",
		"много", "немного", "мало", "несколько",

		// Временные указатели
		"сейчас", "теперь", "сегодня", "вчера", "завтра",
		"потом", " затем", "после", "всегда", "никогда",
		"иногда", "часто", "редко",

		// Пространственные указатели
		"здесь", "тут", "там", "везде", "всюду", "нигде",
		"вверху", "внизу", "справа", "слева", "впереди",
		"сзади", "внутри", "снаружи", "близко", "далеко",

		// Другие частотные слова
		"вот", "всё", "все", "весь", "вся", "всё", "всех",
		"каждый", "каждая", "каждое", "каждые", "любой",
		"любая", "любое", "любые", "некоторый", "некоторая",
		"некоторое", "некоторые", "некий", "некая", "некое",
		"некие", "какой-то", "какая-то", "какое-то", "какие-то",
		"чей-то", "чья-то", "чьё-то", "чьи-то",

		// Модальные слова
		"можно", "нужно", "надо", "необходимо", "должен",
		"должна", "должно", "должны", "возможно", "вероятно",
		"наверное", "может", "может быть",

		// Служебные слова для текста
		"например", "так", "итак", "потом", " затем", "кстати",
		"вообще", "конечно", "действительно", "правда",
		"собственно", "значит", "таким", "образом",
	})
	// Active the Spanish.
	language.SetActiveLanguage("ru")

	// Add text.
	tr.Populate(rawText, language, rule)
	// Run the ranking.
	tr.Ranking(algorithmDef)

	// Get all phrases order by weight.
	rankedPhrases := textrank.FindPhrases(tr)
	// Most important phrase.
	fmt.Println(rankedPhrases[0])

	// Get all words order by weight.
	words := textrank.FindSingleWords(tr)
	// Most important word.
	fmt.Println(words[0])

	// Get the most important 10 sentences. Importance by phrase weights.
	sentences := textrank.FindSentencesByRelationWeight(tr, 10)
	// Found sentences
	fmt.Println(sentences)

	// Get the most important 10 sentences. Importance by word occurrence.
	sentences = textrank.FindSentencesByWordQtyWeight(tr, 10)
	// Found sentences
	fmt.Println(sentences)

	// Get the first 10 sentences, start from 5th sentence.
	sentences = textrank.FindSentencesFrom(tr, 5, 10)
	// Found sentences
	fmt.Println(sentences)

	// Get sentences by phrase/word chains order by position in text.
	sentencesPh := textrank.FindSentencesByPhraseChain(tr, []string{"gnome", "shell", "extension"})
	// Found sentence.
	fmt.Println(sentencesPh[0])
}
