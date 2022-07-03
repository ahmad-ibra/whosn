package models

import (
	"time"
)

// EventUser holds data coming from the event_users table
type EventUser struct {
	EventID   string    `json:"event_id"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (eventUser *EventUser) ConstructCreate(eventID string, userID string) {
	eventUser.EventID = eventID
	eventUser.UserID = userID
}
