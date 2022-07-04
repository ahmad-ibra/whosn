package models

import (
	"time"

	"github.com/Ahmad-Ibra/whosn-core/internal/config"
)

// Event holds data coming from the events table
type Event struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	OwnerID   string    `json:"owner_id"`
	Time      time.Time `json:"time"`
	Location  string    `json:"location"`
	MinUsers  uint64    `json:"min_users"`
	MaxUsers  uint64    `json:"max_users"`
	Price     float64   `json:"price"`
	Link      string    `json:"link"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (event *Event) ConstructCreate(ownerID string) {
	event.OwnerID = ownerID
	event.Link = generateEventLink()
}

func (event *Event) ConstructUpdate(original *Event) {
	event.ID = original.ID
	if event.Name == "" {
		event.Name = original.Name
	}
	event.OwnerID = original.OwnerID
	if event.Time.IsZero() {
		event.Time = original.Time
	}
	if event.Location == "" {
		event.Location = original.Location
	}
	if event.MinUsers == 0 {
		event.MinUsers = original.MinUsers
	}
	if event.MaxUsers == 0 {
		event.MaxUsers = original.MaxUsers
	}
	if event.Price == 0 {
		event.Price = original.Price
	}
	event.Link = original.Link
	event.CreatedAt = original.CreatedAt
	event.UpdatedAt = time.Now().UTC()
}

func generateEventLink() string {
	cfg := config.GetConfig()
	return "http://" + cfg.FrontendDomain + "/event/"
}
