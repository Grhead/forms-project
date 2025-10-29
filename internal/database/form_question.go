package database

type dbFormsQuestion struct {
	Id         string       `gorm:"primaryKey"`
	FormId     []dbForm     `gorm:"foreignKey:Id"`
	QuestionId []dbQuestion `gorm:"foreignKey:Id"`
}
