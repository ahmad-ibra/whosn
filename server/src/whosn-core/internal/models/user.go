package models

import (
	"time"
)

// User holds data coming from the users table
type User struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Username    string    `json:"user_name"`
	Password    string    `json:"password"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
