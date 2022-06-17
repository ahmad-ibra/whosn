package models

// DataStorer is an interface for our datasource
type DataStorer interface {
	ListAllEvents() (*[]Event, error)
	GetEventByID(eventID string) (*Event, error)
	InsertEvent(event Event) error
	UpdateEventByID(eventUpdate Event, eventID string) (*Event, error)
	DeleteEventByID(eventID string) error
	ListAllUsers() (*[]User, error)
	GetUserByID(userID string) (*User, error)
	GetUserByUsername(username string) (*User, error)
	InsertUser(user User) error
	UpdateUserByID(userUpdate User, userID string) (*User, error)
	DeleteUserByID(userID string) error
}
