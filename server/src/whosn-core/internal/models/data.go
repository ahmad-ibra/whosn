package models

import (
	"errors"
	"sync"
	"time"
)

var (
	// mock events till we get a db in place
	events = []Event{
		{
			ID:         "f503857c-5334-450d-be87-15bdcde50341",
			Name:       "Volleyball",
			StartTime:  time.Time{},
			Location:   "6Pack",
			MinUsers:   10,
			MaxUsers:   12,
			Price:      120.00,
			IsFlatRate: false,
			OwnerID:    "f503857c-5334-450d-be87-15bdcde50342",
			Link:       "www.somepage.com/abasdcasdfasdf/1",
			CreatedAt:  time.Time{},
			UpdatedAt:  time.Time{},
		},
		{
			ID:         "50262b10-3d8e-4134-9869-1e0ed5cfe9f7",
			Name:       "Soccer",
			StartTime:  time.Time{},
			Location:   "Tom binnie",
			MinUsers:   10,
			MaxUsers:   22,
			Price:      155.00,
			IsFlatRate: false,
			OwnerID:    "f503857c-5334-450d-be87-15bdcde50343",
			Link:       "www.somepage.com/abasdcasdfasdf/2",
			CreatedAt:  time.Time{},
			UpdatedAt:  time.Time{},
		},
		{
			ID:         "45de396a-4880-4c52-9689-f8812bf67a51",
			Name:       "Movie",
			StartTime:  time.Time{},
			Location:   "Landmarks Guildford",
			MinUsers:   1,
			MaxUsers:   10,
			Price:      12,
			IsFlatRate: true,
			OwnerID:    "f503857c-5334-450d-be87-15bdcde50344",
			Link:       "www.somepage.com/abasdcasdfasdf/3",
			CreatedAt:  time.Time{},
			UpdatedAt:  time.Time{},
		},
	}

	// mock users till we get a db in place
	users = []User{
		{
			ID:          "7076f342-fd08-4d44-a7ca-baeb31e581fe",
			Name:        "Ahmad I",
			Username:    "aibra",
			Password:    "abc123",
			Email:       "email1@whosn.xyz.com",
			PhoneNumber: "604-534-6333",
			CreatedAt:   time.Time{},
			UpdatedAt:   time.Time{},
		},
		{
			ID:          "b1be816f-fb34-4ab4-a1de-d3a08eca5217",
			Name:        "Karrar A",
			Username:    "karol-a",
			Password:    "qwerty",
			Email:       "email23234234@whosn.xyz.com",
			PhoneNumber: "778-111-6333",
			CreatedAt:   time.Time{},
			UpdatedAt:   time.Time{},
		},
		{
			ID:          "489c800e-034b-4225-bfb1-3327652b63cb",
			Name:        "Wael A",
			Username:    "waelus-ice-wizard",
			Password:    "999888777",
			Email:       "anotherEmail@whosn.xyz.com",
			PhoneNumber: "123-345-4567",
			CreatedAt:   time.Time{},
			UpdatedAt:   time.Time{},
		},
	}
)

// Storer is an interface for our datasource
type Storer interface {
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

// DataStore holds a database connection
type DataStore struct{}

// Compile time check that DataStore implements the Storer interface
var _ Storer = (*DataStore)(nil)

var lock = &sync.Mutex{}
var dataStore *DataStore

func GetDataStore() *DataStore {
	lock.Lock()
	defer lock.Unlock()
	if dataStore == nil {
		dataStore = &DataStore{}
	}
	return dataStore
}

// TODO: Ensure all db changing functions are thread safe

// ListAllEvents gets every event in our datasource
func (d *DataStore) ListAllEvents() (*[]Event, error) {
	return &events, nil
}

// GetEventByID gets a single event in our datasource
func (d *DataStore) GetEventByID(eventID string) (*Event, error) {
	for _, event := range events {
		if event.ID == eventID {
			return &event, nil
		}
	}
	return nil, errors.New("event not found")
}

// InsertEvent inserts the event into the datasource
func (d *DataStore) InsertEvent(event Event) error {
	events = append(events, event)
	return nil
}

// UpdateEventByID updates the event in the datasource
func (d *DataStore) UpdateEventByID(eventUpdate Event, eventID string) (*Event, error) {
	for i := 0; i < len(events); i++ {
		event := &events[i]
		if event.ID == eventID {
			event.UpdatedAt = time.Now()
			if eventUpdate.Name != "" {
				event.Name = eventUpdate.Name
			}
			if !eventUpdate.StartTime.IsZero() {
				event.StartTime = eventUpdate.StartTime
			}
			if eventUpdate.Location != "" {
				event.Location = eventUpdate.Location
			}
			// TODO: dont allow setting MinUsers above MaxUsers
			if eventUpdate.MinUsers != 0 {
				event.MinUsers = eventUpdate.MinUsers
			}
			if eventUpdate.MaxUsers != 0 {
				event.MaxUsers = eventUpdate.MaxUsers
			}
			if eventUpdate.Price != 0 {
				event.Price = eventUpdate.Price
			}
			// Note: frontend needs to make sure that its always passing this value through
			if eventUpdate.IsFlatRate != event.IsFlatRate {
				event.IsFlatRate = eventUpdate.IsFlatRate
			}
			return event, nil
		}
	}
	return nil, errors.New("event not found")
}

// DeleteEventByID deletes the event in the datasource
func (d *DataStore) DeleteEventByID(eventID string) error {
	for i, event := range events {
		if event.ID == eventID {
			events = append(events[:i], events[i+1:]...)
			return nil
		}
	}
	return errors.New("event not found")
}

// ListAllUsers gets every event in our datasource
func (d *DataStore) ListAllUsers() (*[]User, error) {
	return &users, nil
}

// GetUserByID gets a single user in our datasource
func (d *DataStore) GetUserByID(userID string) (*User, error) {
	for _, user := range users {
		if user.ID == userID {
			return &user, nil
		}
	}
	return nil, errors.New("user not found")
}

// GetUserByUsername gets a single user in our datasource
func (d *DataStore) GetUserByUsername(username string) (*User, error) {
	for _, user := range users {
		if user.Username == username {
			return &user, nil
		}
	}
	return nil, errors.New("user not found")
}

// InsertUser inserts the user into the datasource
func (d *DataStore) InsertUser(user User) error {
	users = append(users, user)
	return nil
}

// UpdateUserByID updates the user in the datasource
func (d *DataStore) UpdateUserByID(userUpdate User, userID string) (*User, error) {
	for i := 0; i < len(users); i++ {
		user := &users[i]
		if user.ID == userID {
			user.UpdatedAt = time.Now()
			if userUpdate.Name != "" {
				user.Name = userUpdate.Name
			}
			if userUpdate.Email != "" {
				user.Email = userUpdate.Email
			}
			if userUpdate.PhoneNumber != "" {
				user.PhoneNumber = userUpdate.PhoneNumber
			}
			if userUpdate.Username != "" {
				user.Username = userUpdate.Username
			}
			if userUpdate.Password != "" {
				user.Password = userUpdate.Password
			}
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}

// DeleteUserByID deletes the user in the datasource
func (d *DataStore) DeleteUserByID(userID string) error {
	for i, user := range users {
		if user.ID == userID {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return errors.New("user not found")
}
