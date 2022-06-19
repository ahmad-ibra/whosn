package models

import (
	"time"

	"github.com/Ahmad-Ibra/whosn-core/internal/config"

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
	event.Link = generateEventLink(event.ID)
}

func generateEventLink(eventID string) string {
	cfg := config.GetConfig()
	return "https://" + cfg.FrontendDomain + "/event/" + eventID
}
