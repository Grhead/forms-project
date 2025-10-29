package app

import (
	"log"
	"tusur-forms/internal/config"
	"tusur-forms/internal/database"
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
	db, err := dbProvider.NewDbConfig("C:\\Users\\Yomayo\\GolandProjects\\forms-project\\local\\forms.db")
	log.Println("Successfully connected to database")
	if err != nil {
		return err
	}
	err = database.Migrate(db)
	log.Println("Successfully migrated database")
	if err != nil {
		return err
	}
	return nil
}
