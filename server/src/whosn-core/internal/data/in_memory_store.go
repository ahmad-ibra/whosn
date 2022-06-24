package data

import (
	"net/http"
	"sync"
	"time"

	"github.com/Ahmad-Ibra/whosn-core/internal/data/models"
	wnerr "github.com/Ahmad-Ibra/whosn-core/internal/errors"
)

var (
	// mock events till we get a db in place
	events []models.Event

	// mock users till we get a db in place
	users []models.User

	// mock eventUsers till we get a db in place
	eventUsers []models.EventUser
)

type inMemoryStore struct{}

// Compile time check that DataStore implements the Storer interface
var _ Storer = (*inMemoryStore)(nil)

var lock = &sync.Mutex{}
var dataStore *inMemoryStore

func GetInMemoryStore() *inMemoryStore {
	lock.Lock()
	defer lock.Unlock()
	if dataStore == nil {
		dataStore = &inMemoryStore{}
	}
	return dataStore
}

// TODO: Ensure all db changing functions are thread safe

// ListAllEvents gets every event in our datasource
func (d *inMemoryStore) ListAllEvents() (*[]models.Event, error) {
	return &events, nil
}

// ListJoinedEvents gets every joined event in our datasource
func (d *inMemoryStore) ListJoinedEvents(userID string) (*[]models.Event, error) {
	var joinedEvents []models.Event

	for _, event := range events {
		_, err := d.GetEventUserByEventIDUserID(event.ID, userID)
		if err != nil {
			if err, ok := err.(*wnerr.WnError); ok && err.StatusCode == http.StatusNotFound {
				continue
			}
			return &[]models.Event{}, err
		}
		joinedEvents = append(joinedEvents, event)
	}
	return &joinedEvents, nil
}

// ListOwnedEvents gets every owned event in our datasource
func (d *inMemoryStore) ListOwnedEvents(userId string) (*[]models.Event, error) {
	var ownedEvents []models.Event
	for _, event := range events {
		if event.OwnerID == userId {
			ownedEvents = append(ownedEvents, event)
		}
	}
	return &ownedEvents, nil
}

// GetEventByID gets a single event in our datasource
func (d *inMemoryStore) GetEventByID(eventID string) (*models.Event, error) {
	for _, event := range events {
		if event.ID == eventID {
			return &event, nil
		}
	}
	return &models.Event{}, wnerr.NewError(http.StatusNotFound, "event not found")
}

// InsertEvent inserts the event into the datasource
func (d *inMemoryStore) InsertEvent(event models.Event) error {
	events = append(events, event)
	return nil
}

// UpdateEventByID updates the event in the datasource
func (d *inMemoryStore) UpdateEventByID(eventUpdate models.Event, eventID string) (*models.Event, error) {
	for i := 0; i < len(events); i++ {
		event := &events[i]
		if event.ID == eventID {
			event.UpdatedAt = time.Now()
			if eventUpdate.Name != "" {
				event.Name = eventUpdate.Name
			}
			if eventUpdate.Time != "" {
				event.Time = eventUpdate.Time
			}
			if eventUpdate.Location != "" {
				event.Location = eventUpdate.Location
			}
			if eventUpdate.MinUsers != 0 {
				event.MinUsers = eventUpdate.MinUsers
			}
			if eventUpdate.MaxUsers != 0 {
				event.MaxUsers = eventUpdate.MaxUsers
			}
			if eventUpdate.Price != 0 {
				event.Price = eventUpdate.Price
			}
			return event, nil
		}
	}
	return &models.Event{}, wnerr.NewError(http.StatusNotFound, "event not found")
}

// DeleteEventByID deletes the event in the datasource
func (d *inMemoryStore) DeleteEventByID(eventID string) error {
	for i, event := range events {
		if event.ID == eventID {
			events = append(events[:i], events[i+1:]...)
			return nil
		}
	}
	return wnerr.NewError(http.StatusNotFound, "event not found")
}

// ListAllUsers gets every event in our datasource
func (d *inMemoryStore) ListAllUsers() (*[]models.User, error) {
	return &users, nil
}

// GetUserByID gets a single user in our datasource
func (d *inMemoryStore) GetUserByID(userID string) (*models.User, error) {
	for _, user := range users {
		if user.ID == userID {
			return &user, nil
		}
	}
	return &models.User{}, wnerr.NewError(http.StatusNotFound, "user not found")
}

// GetUserByUsername gets a single user in our datasource
func (d *inMemoryStore) GetUserByUsername(username string) (*models.User, error) {
	for _, user := range users {
		if user.Username == username {
			return &user, nil
		}
	}
	return &models.User{}, wnerr.NewError(http.StatusNotFound, "user not found")
}

// InsertUser inserts the user into the datasource
func (d *inMemoryStore) InsertUser(user models.User) error {
	users = append(users, user)
	return nil
}

// UpdateUserByID updates the user in the datasource
func (d *inMemoryStore) UpdateUserByID(userUpdate models.User, userID string) (*models.User, error) {
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
	return &models.User{}, wnerr.NewError(http.StatusNotFound, "user not found")
}

// DeleteUserByID deletes the user in the datasource
func (d *inMemoryStore) DeleteUserByID(userID string) error {
	for i, user := range users {
		if user.ID == userID {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return wnerr.NewError(http.StatusNotFound, "user not found")
}

// ListAllEventUsers gets every eventUser in our datasource
func (d *inMemoryStore) ListAllEventUsers() (*[]models.EventUser, error) {
	return &eventUsers, nil
}

// GetEventUserByEventIDUserID gets a single eventUser in our datasource
func (d *inMemoryStore) GetEventUserByEventIDUserID(eventID string, userID string) (*models.EventUser, error) {
	for _, eventUser := range eventUsers {
		if eventUser.EventID == eventID && eventUser.UserID == userID {
			return &eventUser, nil
		}
	}
	return &models.EventUser{}, wnerr.NewError(http.StatusNotFound, "eventUser not found")
}

// InsertEventUser inserts the eventUser into the datasource
func (d *inMemoryStore) InsertEventUser(eventUser models.EventUser) error {
	eventUsers = append(eventUsers, eventUser)
	return nil
}

// DeleteEventUserByID deletes the eventUser in the datasource
func (d *inMemoryStore) DeleteEventUserByID(eventUserID string) error {
	for i, eventUser := range eventUsers {
		if eventUser.ID == eventUserID {
			eventUsers = append(eventUsers[:i], eventUsers[i+1:]...)
			return nil
		}
	}
	return wnerr.NewError(http.StatusNotFound, "eventUser not found")
}
