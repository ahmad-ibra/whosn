package data

import (
	"github.com/Ahmad-Ibra/whosn-core/internal/data/models"
)

// Storer is an interface for our datasource
type Storer interface {
	ListAllUsers() (*[]models.User, error)
	GetUserByID(userID string) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	InsertUser(user models.User) error
	UpdateUserByID(userUpdate models.User, userID string) (*models.User, error)
	DeleteUserByID(userID string) error

	ListAllEvents() (*[]models.Event, error)
	ListJoinedEvents(userID string) (*[]models.Event, error)
	ListOwnedEvents(userID string) (*[]models.Event, error)
	GetEventByID(eventID string) (*models.Event, error)
	InsertEvent(event models.Event) error
	UpdateEventByID(eventUpdate models.Event, eventID string) (*models.Event, error)
	DeleteEventByID(eventID string) error

	ListAllEventUsers() (*[]models.EventUser, error)
	GetEventUserByEventIDUserID(eventID string, userID string) (*models.EventUser, error)
	InsertEventUser(eventUser models.EventUser) error
	DeleteEventUserByID(eventUserID string) error
}
