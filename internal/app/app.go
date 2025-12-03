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
	"tusur-forms/internal/transport"

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
	transportEntity := transport.NewOrchestrator(newOrchestrator)

	r.Get("/questions", transportEntity.GetQuestions)
	r.Post("/question", transportEntity.CreateQuestion)

	r.Get("/form", transportEntity.GetForm)
	r.Get("/forms", transportEntity.GetForms)
	r.Post("/form", transportEntity.CreateForm)

	r.Post("/generate", transportEntity.GenerateXlsx)

	fmt.Println("Server started on http://localhost:3000")
	err = http.ListenAndServe(":3000", r)
	if err != nil {
		return err
	}
	//f, err := newOrchestrator.CheckoutAnswers("ae66e57e-b0dd-4404-836c-9c5d015f0309")
	//if err != nil {
	//	return err
	//}
	//log.Println(f.Print())
	//
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
	//err = file.SaveFile(index, "test.xlsx")
	//if err != nil {
	//	return err
	//}
	return nil
}
