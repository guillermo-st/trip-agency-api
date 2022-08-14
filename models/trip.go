package models

import "time"

type Trip struct {
	Id         uint64     `json:"id" db:"trip_id"`
	UserId     uint64     `json:"user_id" db:"user_id"`
	StartedAt  time.Time  `json:"started_at" db:"started_at"`
	FinishedAt *time.Time `json:"finished_at" db:"finished_at"`
}
