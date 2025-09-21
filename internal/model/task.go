package model

import "time"

type Task struct {
	ID          int
	Description string
	Status      string // todo done in-progress
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
