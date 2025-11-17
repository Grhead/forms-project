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

	quest1 := domain.Question{
		Title:       "Это первый вопрос",
		Description: "Ответь на любой RADIO",
		Type: domain.QuestionType{
			Title: domain.TypeRadio,
		},
		IsRequired:      true,
		PossibleAnswers: []*domain.PossibleAnswer{{Content: "Первый RADIO"}, {Content: "Второй RADIO"}, {Content: "Третий RADIO"}, {Content: "Четвёртый RADIO"}},
	}
	quest2 := domain.Question{
		Title:       "Это второй вопрос",
		Description: "Это оставь пустым",
		Type: domain.QuestionType{
			Title: domain.TypeText,
		},
		IsRequired:      true,
		PossibleAnswers: []*domain.PossibleAnswer{},
	}
	quest3 := domain.Question{
		Title:       "Это третий вопрос",
		Description: "Ответь на любой RADIO",
		Type: domain.QuestionType{
			Title: domain.TypeRadio,
		},
		IsRequired:      true,
		PossibleAnswers: []*domain.PossibleAnswer{{Content: "Первый-второй RADIO"}, {Content: "Второй-третий RADIO"}, {Content: "Третий-пятый RADIO"}, {Content: "Четвёртый RADIO"}},
	}
	quest4 := domain.Question{
		Title:       "Это четвертый вопрос",
		Description: "Ответь на любой RADIO",
		Type: domain.QuestionType{
			Title: domain.TypeRadio,
		},
		IsRequired:      true,
		PossibleAnswers: []*domain.PossibleAnswer{{Content: "Первый-второй RADIO"}, {Content: "Не второй RADIO"}, {Content: "Не третий RADIO"}, {Content: "Не четвёртый RADIO"}},
	}
	quest5 := domain.Question{
		Title:       "Это пятый вопрос",
		Description: "Ответь следующим текстом (просто скопируй): 'Высокая громкость вредит вашему слуху'",
		Type: domain.QuestionType{
			Title: domain.TypeCheckbox,
		},
		IsRequired:      false,
		PossibleAnswers: []*domain.PossibleAnswer{{Content: "Первый CHECK"}, {Content: "Второй CHECK"}, {Content: "Третий CHECK"}, {Content: "Четвёртый CHECK"}},
	}
	quest6 := domain.Question{
		Title:       "Это шестой вопрос",
		Description: "Оставь пустым, не трогай",
		Type: domain.QuestionType{
			Title: domain.TypeText,
		},
		IsRequired:      false,
		PossibleAnswers: []*domain.PossibleAnswer{},
	}
	quest7 := domain.Question{
		Title:       "Это седьмой вопрос",
		Description: "Отметить более 3 чекбоксов",
		Type: domain.QuestionType{
			Title: domain.TypeCheckbox,
		},
		IsRequired:      true,
		PossibleAnswers: []*domain.PossibleAnswer{{Content: "CHECK GANG"}, {Content: "Отметь меня"}, {Content: "И меня отметь"}, {Content: "А меня как хочешь"}},
	}
	forma, err := newOrchestrator.CheckoutForm("Testing ради Answers", "Попытка №17", []*domain.Question{&quest1, &quest2, &quest3, &quest4, &quest5, &quest6, &quest7})
	if err != nil {
		return err
	}
	f, err := gormRepo.GetForm(forma.ID)
	if err != nil {
		return err
	}
	fmt.Println(f.Print())
	return nil
}
