package app

import (
	"log"
	"time"
	"tusur-forms/internal/config"
	"tusur-forms/internal/database"
	"tusur-forms/internal/domain"
)

func Run() error {
	//ctx := context.Background()
	//formProvider := &config.EnvProvider{}
	//cfg, err := formProvider.NewFormConfig()
	//if err != nil {
	//	return err
	//}
	dbProvider := &config.DbSQLiteProvider{}
	log.Println("Connecting to database ...")
	db, err := dbProvider.NewDbConfig("C:\\Users\\egorm\\GolandProjects\\tusur-forms\\local\\forms.db")
	if err != nil {
		return err
	}
	log.Println("Successfully connected to database")
	//err = database.Migrate(db)
	//if err != nil {
	//	return err
	//}
	//log.Println("Successfully migrated database")
	a := &domain.Answer{
		Id:          "1",
		SubmittedAt: time.Now(),
		Content:     "Horns",
		FormId:      "1",
		QuestionId:  "1",
	}
	err = database.CreateAnswer(a, db)
	if err != nil {
		return err
	}
	//t := &domain.QuestionType{
	//	Id:    "1",
	//	Title: "Первый тип",
	//}
	//database.CreateQuestionType(t, db)
	//q := &domain.Question{
	//	Id:              "1",
	//	Title:           "Вопрос №1",
	//	Type:            *t,
	//	IsRequired:      true,
	//	PossibleAnswers: nil,
	//}
	//database.CreateQuestion(q, db)
	return nil
}
