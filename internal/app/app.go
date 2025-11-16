package app

import (
	"context"
	"fmt"
	"log"
	"tusur-forms/internal/config"
	"tusur-forms/internal/database"
	"tusur-forms/internal/domain"
	"tusur-forms/internal/services/google"
	"tusur-forms/internal/services/orchectrators"

	"github.com/google/uuid"
)

func Run() error {
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
	tokenConfig, err := config.ReadToken(formCfg.TokenPath)
	if err != nil {
		return err
	}
	tokenSource := oauthConfig.TokenSource(ctx, tokenConfig)

	googleProvider := google.GoogleForms{
		TokenSource: tokenSource,
	}
	dbProvider := &config.DBSQLiteProvider{}
	db, err := dbProvider.Connect(dbCfg)
	if err != nil {
		return err
	}
	gormRepo := database.NewGormRepository(db)
	exists, err := gormRepo.CheckExists()
	if err != nil {
		return err
	}

	if !exists {
		err = gormRepo.Migrate()
		if err != nil {
			return err
		}
		log.Println("Successfully migrated database")
	}

	service, err := googleProvider.NewService(ctx, gormRepo)
	if err != nil {
		return err
	}
	newOrchestrator := orchectrators.NewFormsOrchestrator(service, gormRepo)

	quest := domain.Question{
		ID:          uuid.NewString(),
		Title:       "Simple Question",
		Description: "Give me answers",
		Type: domain.QuestionType{
			ID:    uuid.NewString(),
			Title: "RADIO",
		},
		IsRequired:      true,
		PossibleAnswers: []*domain.PossibleAnswer{{Content: "First answer of universe"}, {Content: "Second answer of Earth"}},
	}
	_, err = newOrchestrator.CheckoutForm("Testing уже не на паре", "Testing", []*domain.Question{&quest})
	if err != nil {
		return err
	}
	f, err := gormRepo.GetForm("8a227a28-32f1-49f8-8867-0f371ca3743e")
	if err != nil {
		return err
	}
	fmt.Println(f.Print())
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
