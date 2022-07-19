package models

import (
	"time"
)

// EventUser holds data coming from the event_users table
type EventUser struct {
	EventID   string    `json:"event_id"`
	UserID    string    `json:"user_id"`
	HasPaid   bool      `json:"has_paid"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// SetPaidBody holds the body of the request to set if a user has paid or not
type SetPaidBody struct {
	HasPaid bool `json:"has_paid"`
}

func (eventUser *EventUser) ConstructCreate(eventID string, userID string) {
	eventUser.EventID = eventID
	eventUser.UserID = userID
}
