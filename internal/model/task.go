package model

import "time"

// Task represents single task in the system
type Task struct {
	ID          int
	Description string
	Status      string // todo done in-progress
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
