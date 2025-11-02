package database

import "gorm.io/gorm"

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&dbForm{},
		&dbQuestionType{},
		&dbQuestion{},
		&dbAnswer{},
		&dbFormsQuestion{},
		&dbPossibleAnswer{},
		&dbQuestionPossibleAnswer{})
	return err
}
