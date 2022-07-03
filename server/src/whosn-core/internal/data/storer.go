package data

import (
	"github.com/Ahmad-Ibra/whosn-core/internal/data/models"
)

// Storer is an interface for our datasource
type Storer interface {
	GetUserByID(userID string) (*models.User, error)
	GetUserByUserName(userName string) (*models.User, error)
	InsertUser(user *models.User) error
	UpdateUserByID(user *models.User, userID string) error
	DeleteUserByID(userID string) error

	ListJoinedEvents(userID string) (*[]models.Event, error)
	ListOwnedEvents(userID string) (*[]models.Event, error)
	GetEventByID(eventID string) (*models.Event, error)
	InsertEvent(event *models.Event) error
	UpdateEventByID(event *models.Event, eventID string) error
	DeleteEventByID(eventID string) error

	GetEventUserByEventIDUserID(eventID string, userID string) (*models.EventUser, error)
	InsertEventUser(eventUser *models.EventUser) error
	DeleteEventUserByEventIDUserID(eventID string, userID string) error
}
