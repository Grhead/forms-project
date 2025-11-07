package app

import (
	"context"
	"log"
	"tusur-forms/internal/config"
	"tusur-forms/internal/database"
	"tusur-forms/internal/services/forms"
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
	googleProvider := services.GoogleForms{
		OauthCfg: oauthConfig,
	}
	service, err := googleProvider.NewService(ctx, "token.json")
	if err != nil {
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
