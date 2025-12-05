package app

import (
	"context"
	"log"
	"net/http"
	"tusur-forms/internal/config"
	"tusur-forms/internal/repository"
	"tusur-forms/internal/services/forms/google"
	"tusur-forms/internal/services/orchectrators"
	"tusur-forms/internal/transport"

	_ "tusur-forms/docs"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger"
)

// Run
// @title tusur-forms-api
// @version 1.0
// @description This is tusur-forms-api
// @termsOfService http://swagger.io/terms/
// @host localhost:3000
// @schemes http
// @BasePath /api
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
	transportEntity := transport.NewOrchestrator(newOrchestrator)
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write(html)
	})
	r.Get("/swagger/*", httpSwagger.WrapHandler)
	r.Route("/api", func(r chi.Router) {
		r.Get("/questions", transportEntity.GetQuestions)
		r.Post("/question", transportEntity.CreateQuestion)

		r.Get("/form", transportEntity.GetForm)
		r.Get("/forms", transportEntity.GetForms)
		r.Post("/form", transportEntity.CreateForm)

		r.Post("/generate", transportEntity.GenerateXlsx)
	})

	//form, err := newOrchestrator.CheckoutAnswers("7df1df32-c91c-4cb3-ac62-8d04e3c6aa89")
	//if err != nil {
	//	return err
	//}
	//file := reports.CreateFile()
	//index, err := file.CreateSpreadsheet(form.Title)
	//if err != nil {
	//	return err
	//}
	//err = file.SetHeader(index, form)
	//if err != nil {
	//	return err
	//}
	//err = file.SetData(index, form)
	//if err != nil {
	//	return err
	//}
	//err = file.SetMediumQuestion(index, form)
	//if err != nil {
	//	return err
	//}
	//err = file.SetMediumDiscipline(index, form)
	//if err != nil {
	//	return err
	//}
	//
	//err = file.SaveFile(index, "test.xlsx")
	//if err != nil {
	//	return err
	//}
	log.Println("Server started on http://localhost:3000")
	err = http.ListenAndServe(":3000", r)
	if err != nil {
		return err
	}
	return nil
}
