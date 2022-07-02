package data

import (
	"fmt"

	"github.com/Ahmad-Ibra/whosn-core/internal/config"
	"github.com/go-pg/migrations/v8"
	"github.com/go-pg/pg/v10"
)

func NewDB() (*pg.DB, error) {
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
		fmt.Printf("migrated from version %v to %v\n", oldV, newV)
	} else {
		fmt.Printf("on version %v\n", oldV)
	}
	// return the db connection
	return db, nil
}
