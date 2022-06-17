package data

import (
	"github.com/Ahmad-Ibra/whosn-core/internal/data/models"
)

// Storer is an interface for our datasource
type Storer interface {
	ListAllEvents() (*[]models.Event, error)
	GetEventByID(eventID string) (*models.Event, error)
	InsertEvent(event models.Event) error
	UpdateEventByID(eventUpdate models.Event, eventID string) (*models.Event, error)
	DeleteEventByID(eventID string) error
	ListAllUsers() (*[]models.User, error)
	GetUserByID(userID string) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	InsertUser(user models.User) error
	UpdateUserByID(userUpdate models.User, userID string) (*models.User, error)
	DeleteUserByID(userID string) error
}
