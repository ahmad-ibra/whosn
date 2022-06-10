package data

import (
	"time"
)

// Group holds data coming from the groups table
type Group struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	OwnerID   uint64    `json:"owner_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Event holds data coming from the events table
type Event struct {
	ID        uint64    `json:"id"`
	GroupID   uint64    `json:"group_id"`
	Name      string    `json:"name"`
	StartTime time.Time `json:"start_time"`
	Location  string    `json:"location"`
	MinUsers  uint64    `json:"min_users"`
	MaxUsers  uint64    `json:"max_users"`
	Price     float64   `json:"price"`
	OwnerID   uint64    `json:"owner_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// User holds data coming from the users table
type User struct {
	ID          uint64    `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// EventUser holds data coming from the event_users table
type EventUser struct {
	ID         uint64    `json:"id"`
	EventID    uint64    `json:"event_id"`
	UserID     uint64    `json:"user_id"`
	TotalOwed  float64   `json:"total_owed"`
	TotalPayed float64   `json:"total_payed"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// GroupUser holds data coming from the group_users table
type GroupUser struct {
	ID        uint64    `json:"id"`
	GroupID   uint64    `json:"group_id"`
	UserID    uint64    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
