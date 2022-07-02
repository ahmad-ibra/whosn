package data

import (
	"fmt"

	"github.com/Ahmad-Ibra/whosn-core/internal/config"
	"github.com/Ahmad-Ibra/whosn-core/internal/data/models"
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

	// connect to db
	db := pg.Connect(&pg.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.DBHost, cfg.DBPort),
		User:     cfg.DBUser,
		Password: cfg.DBPassword,
		Database: cfg.DBName,
	})

	// run migrations
	collections := migrations.NewCollection()
	err := collections.DiscoverSQLMigrations("migrations")
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
	//TODO implement me
	panic("implement me")
}

func (p PGStore) GetUserByUsername(username string) (*models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (p PGStore) InsertUser(user models.User) (*models.User, error) {
	_, err := p.Conn.Model(&user).Insert()
	if err != nil {
		return nil, err
	}

	insertedUser := &models.User{}
	err = p.Conn.Model(insertedUser).Where("user_name = ?", user.UserName).Select()
	if err != nil {
		return nil, err
	}

	return insertedUser, nil

}

func (p PGStore) UpdateUserByID(userUpdate models.User, userID string) (*models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (p PGStore) DeleteUserByID(userID string) error {
	//TODO implement me
	panic("implement me")
}

func (p PGStore) ListJoinedEvents(userID string) (*[]models.Event, error) {
	//TODO implement me
	panic("implement me")
}

func (p PGStore) ListOwnedEvents(userID string) (*[]models.Event, error) {
	//TODO implement me
	panic("implement me")
}

func (p PGStore) GetEventByID(eventID string) (*models.Event, error) {
	//TODO implement me
	panic("implement me")
}

func (p PGStore) InsertEvent(event models.Event) error {
	//TODO implement me
	panic("implement me")
}

func (p PGStore) UpdateEventByID(eventUpdate models.Event, eventID string) (*models.Event, error) {
	//TODO implement me
	panic("implement me")
}

func (p PGStore) DeleteEventByID(eventID string) error {
	//TODO implement me
	panic("implement me")
}

func (p PGStore) GetEventUserByEventIDUserID(eventID string, userID string) (*models.EventUser, error) {
	//TODO implement me
	panic("implement me")
}

func (p PGStore) InsertEventUser(eventUser models.EventUser) error {
	//TODO implement me
	panic("implement me")
}

func (p PGStore) DeleteEventUserByID(eventUserID string) error {
	//TODO implement me
	panic("implement me")
}