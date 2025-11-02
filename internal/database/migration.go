package database

import (
	"gorm.io/gorm"
)

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

func CheckExists(db *gorm.DB) (bool, error) {
	tables := []string{
		"db_forms",
		"db_question_types",
		"db_questions",
		"db_forms_questions",
		"db_answers",
		"db_possible_answers",
		"db_question_possible_answers"}
	rows, err := db.Table("sqlite_master").
		Where("type = ?", "table").
		Where("name IN ? ", tables).
		Select("name").
		Rows()
	if err != nil {
		return true, err
	}
	defer rows.Close()
	var count uint8
	for rows.Next() {
		count++
	}
	if count != uint8(len(tables)) {
		return false, nil
	}
	return true, nil
}
