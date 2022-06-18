package models

import (
	"time"

	"github.com/google/uuid"
)

// Event holds data coming from the events table
type Event struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	OwnerID   string    `json:"owner_id"`
	StartTime time.Time `json:"start_time"`
	Location  string    `json:"location"`
	MinUsers  uint64    `json:"min_users"`
	MaxUsers  uint64    `json:"max_users"`
	Price     float64   `json:"price"`
	Link      string    `json:"link"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (event *Event) Construct(ownerID string) {
	event.ID = uuid.New().String()

	curTime := time.Now()
	event.CreatedAt = curTime
	event.UpdatedAt = curTime

	event.OwnerID = ownerID

	// TODO: validate that we have the right link
	// link should be something from the frontend that can get data from https://host/api/v1/secured/event/{eventID}
	event.Link = "https://whosn.xyz/event/" + event.ID
}
