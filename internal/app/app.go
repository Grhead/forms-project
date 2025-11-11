package app

import (
	"context"
	"fmt"
	"log"
	"tusur-forms/internal/config"
	"tusur-forms/internal/database"
	"tusur-forms/internal/domain"
	services "tusur-forms/internal/services/forms"

	"github.com/google/uuid"
)

func Run() error {
	const filename = "C:\\Users\\Egor Mishchuk\\GolandProjects\\forms-project\\configs\\token.json"
	ctx := context.Background()

	cfgProvider := &config.EnvConfigProvider{}
	formCfg, err := cfgProvider.LoadFormConfig()
	if err != nil {
		return err
	}
	dbCfg, err := cfgProvider.LoadDBConfig()
	if err != nil {
		return err
	}
	oauthConfig := config.NewOAuth2Config(formCfg)
	tokenConfig, err := config.ReadToken(filename)
	tokenSource := oauthConfig.TokenSource(ctx, tokenConfig)
	if err != nil {
		log.Fatal("After token source")
		return err
	}
	googleProvider := services.GoogleForms{
		TokenSource: tokenSource,
	}
	service, err := googleProvider.NewService(ctx)
	if err != nil {
		log.Fatal("After services")
		return err
	}
	dbProvider := &config.DBSQLiteProvider{}
	db, err := dbProvider.Connect(dbCfg)
	if err != nil {
		return err
	}

	exists, err := database.CheckExists(db)
	if err != nil {
		return err
	}

	if !exists {
		err = database.Migrate(db)
		if err != nil {
			return err
		}
		log.Println("Successfully migrated database")
	}

	form, err := service.NewForm("Testing на паре", "Testing")
	if err != nil {
		return err
	}
	fmt.Println("Form created")
	form, err = service.GetForm(form.ExternalID)
	if err != nil {
		return err
	}
	fmt.Println(form)
	quest := domain.Question{
		ID:          uuid.NewString(),
		Title:       "Simple Question",
		Description: "Give me answers",
		Type: domain.QuestionType{
			ID:    uuid.NewString(),
			Title: "RADIO",
		},
		IsRequired:      true,
		PossibleAnswers: []domain.PossibleAnswer{{Content: "First answer of universe"}, {Content: "Second answer of Earth"}},
	}
	_, err = service.SetQuestions(form, []*domain.Question{&quest})
	if err != nil {
		return err
	}

	// service.GetForm()

	//a := &domain.Answer{
	//	ID:          "1",
	//	SubmittedAt: time.Now(),
	//	Content:     "Horns",
	//	FormID:      "1",
	//	QuestionID:  "1",
	//}
	//err = database.CreateAnswer(a, db)
	//if err != nil {
	//	return err
	//}
	//t := &domain.QuestionType{
	//	ID:    "1",
	//	Title: "Первый тип",
	//}
	//database.CreateQuestionType(t, db)
	//q := &domain.Question{
	//	ID:              "1",
	//	Title:           "Вопрос №1",
	//	Type:            *t,
	//	IsRequired:      true,
	//	PossibleAnswers: nil,
	//}
	//database.CreateQuestion(q, db)
	return nil
}
