package models

import (
	"database/sql"
	"time"
)

type Trip struct {
	Id         uint         `json:"id" db:"trip_id"`
	UserId     uint         `json:"user_id" db:"user_id"`
	StartedAt  time.Time    `json:"started_at" db:"started_at"`
	FinishedAt sql.NullTime `json:"finished_at" db:"finished_at"`
}
