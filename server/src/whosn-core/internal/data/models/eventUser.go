package models

import (
	"time"
)

// EventUser holds data coming from the event_users table
type EventUser struct {
	ID         string    `json:"id"`
	EventID    uint64    `json:"event_id"`
	UserID     uint64    `json:"user_id"`
	TotalOwed  float64   `json:"total_owed"`
	TotalPayed float64   `json:"total_payed"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
