package database

import (
	"time"
)

type dbForm struct {
	FormId         string `gorm:"primaryKey"`
	FormExternalId string
	FormCreatedAt  time.Time
}
