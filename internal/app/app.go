package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"tusur-forms/internal/config"
	"tusur-forms/internal/repository"
	"tusur-forms/internal/services/forms/google"
	"tusur-forms/internal/services/orchectrators"
	"tusur-forms/internal/services/reports"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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
	gormRepo := repository.NewGormRepository(db)
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
	html, err := cfgProvider.LoadMainHtml()
	if err != nil {
		return err
	}
	service, err := googleProvider.NewService(ctx, gormRepo)
	if err != nil {
		return err
	}
	newOrchestrator := orchectrators.NewFormsOrchestrator(service, gormRepo)
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write(html)
	})

	//r.Get("/questions", getTasks)    // Получить все задачи
	//r.Post("/questions", createTask) // Создать новую задачу
	//
	//r.Get("/forms", getTasks)    // Получить все задачи
	//r.Post("/forms", createTask) // Создать новую задачу
	//
	//r.Post("/generate", createTask) // Создать новую задачу

	// 4. Запускаем сервер
	fmt.Println("Сервер запущен на http://localhost:3000")
	err = http.ListenAndServe(":3000", r)
	if err != nil {
		return err
	}
	//quest1 := domain.Question{
	//	Title:       "Это первый вопрос",
	//	Description: "Ответь на любой RADIO",
	//	Type: domain.QuestionType{
	//		Title: domain.TypeRadio,
	//	},
	//	IsRequired:      true,
	//	PossibleAnswers: []*domain.PossibleAnswer{{Content: "Первый RADIO"}, {Content: "Второй RADIO"}, {Content: "Третий RADIO"}, {Content: "Четвёртый RADIO"}},
	//}
	//quest2 := domain.Question{
	//	Title:       "Это второй вопрос",
	//	Description: "Это оставь пустым",
	//	Type: domain.QuestionType{
	//		Title: domain.TypeText,
	//	},
	//	IsRequired:      false,
	//	PossibleAnswers: []*domain.PossibleAnswer{},
	//}
	//quest3 := domain.Question{
	//	Title:       "Это третий вопрос",
	//	Description: "Ответь на любой RADIO",
	//	Type: domain.QuestionType{
	//		Title: domain.TypeRadio,
	//	},
	//	IsRequired:      true,
	//	PossibleAnswers: []*domain.PossibleAnswer{{Content: "Первый-второй RADIO"}, {Content: "Второй-третий RADIO"}, {Content: "Третий-пятый RADIO"}, {Content: "Четвёртый RADIO"}},
	//}
	//_, err = newOrchestrator.CheckoutForm("Testing ради Answers", "Попытка №19", []*domain.Question{&quest1, &quest2, &quest3})
	//if err != nil {
	//	return err
	//}
	f, err := newOrchestrator.CheckoutAnswers("ae66e57e-b0dd-4404-836c-9c5d015f0309")
	if err != nil {
		return err
	}
	log.Println(f.Print())

	form, err := gormRepo.GetForm("ae66e57e-b0dd-4404-836c-9c5d015f0309")
	if err != nil {
		return err
	}
	if form == nil {
		return fmt.Errorf("form does not exists")
	}
	log.Println(form.Print())
	file := reports.CreateFile()
	index, err := file.CreateSpreadsheet(form.Title)
	if err != nil {
		return err
	}
	err = file.SetHeader(index, form)
	if err != nil {
		return err
	}
	err = file.SetData(index, form)
	if err != nil {
		return err
	}
	err = file.SaveFile(index, "test.xlsx")
	if err != nil {
		return err
	}
	return nil
}
