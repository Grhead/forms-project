package database

import (
	"time"
)

type dbForm struct {
	FormId         string
	FormExternalId string
	FormTimestamp  time.Time
}
