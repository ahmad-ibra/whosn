package data

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Ahmad-Ibra/whosn-core/internal/config"
	"github.com/Ahmad-Ibra/whosn-core/internal/data/models"
	wnerr "github.com/Ahmad-Ibra/whosn-core/internal/errors"
	"github.com/go-pg/migrations/v8"
	"github.com/go-pg/pg/v10"
	log "github.com/sirupsen/logrus"
)

type PGStore struct {
	Conn *pg.DB
}

// Compile time check that DataStore implements the Storer interface
var _ Storer = (*PGStore)(nil)

func NewDB() (*PGStore, error) {
	cfg := config.GetConfig()

	var opts *pg.Options
	var err error
	var migrationsDir string

	if cfg.Env == "prod" {
		opts, err = pg.ParseURL(cfg.DBUrl)
		if err != nil {
			return nil, err
		}
		migrationsDir = "/app/server/src/whosn-core/migrations"

	} else {
		opts = &pg.Options{
			Addr:     fmt.Sprintf("%s:%s", cfg.DBHost, cfg.DBPort),
			User:     cfg.DBUser,
			Password: cfg.DBPassword,
			Database: cfg.DBName,
		}
		if cfg.Env == "test" {
			migrationsDir = "../../migrations"
		} else if cfg.Env == "dev" {
			migrationsDir = "migrations"
		}
	}

	// connect to db
	db := pg.Connect(opts)

	// run migrations
	collections := migrations.NewCollection()

	err = collections.DiscoverSQLMigrations(migrationsDir)
	if err != nil {
		return nil, err
	}

	_, _, err = collections.Run(db, "init")
	if err != nil {
		return nil, err
	}

	oldV, newV, err := collections.Run(db, "up")
	if err != nil {
		return nil, err
	}

	if newV != oldV {
		log.Info(fmt.Sprintf("migrated from version %v to %v\n", oldV, newV))
	} else {
		log.Info(fmt.Sprintf("on version %v\n", oldV))
	}

	// return the db connection
	return &PGStore{Conn: db}, nil
}

func (p PGStore) GetUserByID(userID string) (*models.User, error) {
	user := &models.User{}
	err := p.Conn.Model(user).Where("id = ?", userID).Select()
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (p PGStore) GetUserByUserName(userName string) (*models.User, error) {
	user := &models.User{}
	err := p.Conn.Model(user).Where("user_name = ?", userName).Select()
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (p PGStore) InsertUser(user *models.User) error {
	_, err := p.Conn.Model(user).Insert()
	return err
}

func (p PGStore) UpdateUserByID(user *models.User, userID string) error {
	_, err := p.Conn.Model(user).Where("id = ?", userID).Update()
	return err
}

func (p PGStore) DeleteUserByID(userID string) error {
	ownedEvents, err := p.ListOwnedEvents(userID)
	if err != nil {
		return err
	}

	for _, event := range *ownedEvents {
		err = p.DeleteEventByID(event.ID)
		if err != nil {
			return err
		}
	}

	user := &models.User{}
	_, err = p.Conn.Model(user).Where("id = ?", userID).Delete()
	return err
}

func (p PGStore) ListJoinedEvents(userID string) (*[]models.Event, error) {
	var eventUsers []models.EventUser
	err := p.Conn.Model(&eventUsers).Where("user_id = ?", userID).Select()
	if err != nil {
		return nil, err
	}

	events := make([]models.Event, 0)
	for _, eu := range eventUsers {
		curEvent, err := p.GetEventByID(eu.EventID)
		if err != nil {
			return nil, err
		}
		events = append(events, *curEvent)
	}

	return &events, nil

}

func (p PGStore) ListOwnedEvents(userID string) (*[]models.Event, error) {
	events := make([]models.Event, 0)
	err := p.Conn.Model(&events).Where("owner_id = ?", userID).Select()
	if err != nil {
		return nil, err
	}
	return &events, nil
}

func (p PGStore) GetEventByID(eventID string) (*models.Event, error) {
	event := &models.Event{}
	err := p.Conn.Model(event).Where("id = ?", eventID).Select()
	if err != nil {
		return nil, err
	}

	return event, nil
}

func (p PGStore) InsertEvent(event *models.Event) error {
	_, err := p.Conn.Model(event).Insert()
	return err
}

func (p PGStore) UpdateEventByID(event *models.Event, eventID string) error {
	_, err := p.Conn.Model(event).Where("id = ?", eventID).Update()
	return err
}

func (p PGStore) DeleteEventByID(eventID string) error {
	err := p.DeleteEventUserByEventID(eventID)
	if err != nil {
		return err
	}

	event := &models.Event{}
	_, err = p.Conn.Model(event).Where("id = ?", eventID).Delete()
	return err
}

func (p PGStore) GetEventUserByEventIDUserID(eventID string, userID string) (*models.EventUser, error) {
	eventUser := &models.EventUser{}
	err := p.Conn.Model(eventUser).Where("event_id = ? AND user_id = ?", eventID, userID).Select()
	if err != nil {
		if err.Error() == "pg: no rows in result set" {
			return nil, wnerr.NewError(http.StatusNotFound, "event not found")
		}
		return nil, err
	}

	return eventUser, nil
}

func (p PGStore) InsertEventUser(eventUser *models.EventUser) error {
	_, err := p.Conn.Model(eventUser).Insert()
	return err
}

func (p PGStore) DeleteEventUserByEventID(eventID string) error {
	eventUser := &models.EventUser{}
	_, err := p.Conn.Model(eventUser).Where("event_id = ?", eventID).Delete()
	return err
}

func (p PGStore) DeleteEventUserByEventIDUserID(eventID string, userID string) error {
	eventUser := &models.EventUser{}
	_, err := p.Conn.Model(eventUser).Where("event_id = ? AND user_id = ?", eventID, userID).Delete()
	return err
}

func (p PGStore) ListEventUsers(eventID string) (*[]models.EventUsersIn, error) {
	// get event so we know max participants
	event, err := p.GetEventByID(eventID)
	if err != nil {
		return nil, err
	}

	// get the eventUsers sorted by created_at time
	var eventUsers []models.EventUser
	err = p.Conn.Model(&eventUsers).Where("event_id = ?", eventID).Order("created_at ASC").Select()
	if err != nil {
		return nil, err
	}

	eventUsersIn := make([]models.EventUsersIn, 0)

	// then we iterate over the eventUsers
	for i, eventUser := range eventUsers {
		// get the user
		user, err := p.GetUserByID(eventUser.UserID)
		if err != nil {
			return nil, err
		}

		eui := &models.EventUsersIn{
			EventID:  eventUser.EventID,
			UserID:   eventUser.UserID,
			JoinedAt: eventUser.CreatedAt,
			Name:     user.Name,
			HasPaid:  eventUser.HasPaid,
		}

		if uint64(i) < event.MaxUsers {
			eui.IsIn = true
		} else {
			eui.IsIn = false
		}

		eventUsersIn = append(eventUsersIn, *eui)
	}
	return &eventUsersIn, nil
}

func (p PGStore) SetPaid(eventID string, userID string, hasPaid bool) error {
	eventUser := models.EventUser{}
	updatedAt := time.Now().UTC()
	_, err := p.Conn.Model(&eventUser).Set("has_paid = ?, updated_at = ?", hasPaid, updatedAt).Where("event_id = ? AND user_id = ?", eventID, userID).Update()
	return err
}
