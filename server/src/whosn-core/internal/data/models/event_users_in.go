package models

import (
	"time"
)

// EventUsersIn holds data joined between EventUsers and Users tables
type EventUsersIn struct {
	EventID  string    `json:"event_id"`
	UserID   string    `json:"user_id"`
	JoinedAt time.Time `json:"joined_at"`
	Name     string    `json:"name"`
	IsIn     bool      `json:"is_in"`
}
