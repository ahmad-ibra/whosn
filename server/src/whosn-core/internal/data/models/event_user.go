package models

import (
	"time"

	"github.com/google/uuid"
)

// EventUser holds data coming from the event_users table
type EventUser struct {
	ID        string    `json:"id"`
	EventID   string    `json:"event_id"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (eventUser *EventUser) Construct(eventID string, userID string) {
	curTime := time.Now()
	eventUser.CreatedAt = curTime
	eventUser.UpdatedAt = curTime
	eventUser.ID = uuid.New().String()
	eventUser.EventID = eventID
	eventUser.UserID = userID
}
