package app

import (
	"context"
	"log"
	"tusur-forms/internal/config"
	"tusur-forms/internal/database"
	"tusur-forms/internal/services/forms/google"
	"tusur-forms/internal/services/orchectrators"
	// "context"
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

	// quest1 := domain.Question{
	// 	Title:       "Это первый вопрос",
	// 	Description: "Ответь на любой RADIO",
	// 	Type: domain.QuestionType{
	// 		Title: domain.TypeRadio,
	// 	},
	// 	IsRequired:      true,
	// 	PossibleAnswers: []*domain.PossibleAnswer{{Content: "Первый RADIO"}, {Content: "Второй RADIO"}, {Content: "Третий RADIO"}, {Content: "Четвёртый RADIO"}},
	// }
	// quest2 := domain.Question{
	// 	Title:       "Это второй вопрос",
	// 	Description: "Это оставь пустым",
	// 	Type: domain.QuestionType{
	// 		Title: domain.TypeText,
	// 	},
	// 	IsRequired:      false,
	// 	PossibleAnswers: []*domain.PossibleAnswer{},
	// }
	// quest3 := domain.Question{
	// 	Title:       "Это третий вопрос",
	// 	Description: "Ответь на любой RADIO",
	// 	Type: domain.QuestionType{
	// 		Title: domain.TypeRadio,
	// 	},
	// 	IsRequired:      true,
	// 	PossibleAnswers: []*domain.PossibleAnswer{{Content: "Первый-второй RADIO"}, {Content: "Второй-третий RADIO"}, {Content: "Третий-пятый RADIO"}, {Content: "Четвёртый RADIO"}},
	// }
	// f, err := newOrchestrator.CheckoutForm("Testing ради Answers", "Попытка №19", []*domain.Question{&quest1, &quest2, &quest3})
	// if err != nil {
	// 	return err
	// }
	f, err := newOrchestrator.CheckoutAnswers("ae4941ed-f51e-4e2e-9d96-9781b8a2a67d")
	if err != nil {
		return err
	}
	log.Println(f.Print())
	return nil
}
