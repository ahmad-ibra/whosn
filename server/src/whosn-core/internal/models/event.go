package models

import (
	"time"
)

// Event holds data coming from the events table
type Event struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	OwnerID    uint64    `json:"owner_id"`
	StartTime  time.Time `json:"start_time"`
	Location   string    `json:"location"`
	MinUsers   uint64    `json:"min_users"`
	MaxUsers   uint64    `json:"max_users"`
	Price      float64   `json:"price"`
	IsFlatRate bool      `json:"is_flat_rate"`
	Link       string    `json:"link"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
