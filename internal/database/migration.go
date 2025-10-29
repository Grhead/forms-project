package database

import "gorm.io/gorm"

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&dbForm{},
		&dbQuestion{},
		&dbAnswer{},
		&dbFormsQuestion{},
		&dbQuestionType{},
		&dbPossibleAnswer{},
		&dbQuestionPossibleAnswer{})
	return err
}
