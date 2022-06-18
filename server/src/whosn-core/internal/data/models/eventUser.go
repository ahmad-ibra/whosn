package models

import (
	"time"

	"github.com/Ahmad-Ibra/whosn-core/internal/data"
	"github.com/google/uuid"
)

var ds = data.GetInMemoryStore()

// EventUser holds data coming from the event_users table
type EventUser struct {
	ID         string    `json:"id"`
	EventID    string    `json:"event_id"`
	UserID     string    `json:"user_id"`
	TotalOwed  float64   `json:"total_owed"`
	TotalPayed float64   `json:"total_payed"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (eventUser *EventUser) Construct(eventID string, userID string) error {
	curTime := time.Now()
	eventUser.CreatedAt = curTime
	eventUser.UpdatedAt = curTime
	eventUser.ID = uuid.New().String()
	eventUser.EventID = eventID
	eventUser.UserID = userID
	eventUser.TotalPayed = 0

	event, err := ds.GetEventByID(eventID)
	if err != nil {
		return err
	}
	eventUser.TotalOwed = event.Price
	return nil
}
