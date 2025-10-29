package database

import (
	"time"
)

type dbForm struct {
	Id         string `gorm:"primaryKey"`
	ExternalId string
	CreatedAt  time.Time
}
