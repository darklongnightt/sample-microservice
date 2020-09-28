package db

import (
	"fmt"
	"log"
	"time"

	"github.com/darklongnightt/microservice/config"

	"github.com/go-pg/pg"
)

// Init ...
func Init(config *config.Config, l *log.Logger) (*pg.DB, error) {
	var db *pg.DB
	connectionTries := 5

	for connectionTries > 0 {
		// Setup db connection
		db = pg.Connect(&pg.Options{
			Addr:     config.DB.Host,
			User:     config.DB.User,
			Password: config.DB.Password,
			Database: config.DB.Database,
		})

		// Check if connection is ok
		if _, err := db.Exec("SELECT 1"); err != nil {
			connectionTries--
			l.Printf("failed to connect to db, tries left: %v\n", connectionTries)
			time.Sleep(5 * time.Second)

			if connectionTries == 0 {
				return nil, fmt.Errorf("failed to connect to db\nreason: %v", err)
			}
		} else {
			break
		}
	}

	// Init all tables
	l.Printf("Connection to db successful on %v\n", config.DB.Host)
	if err := CreateAllTables(db); err != nil {
		return nil, fmt.Errorf("failed to init tables\nreason: %v", err)
	}

	return db, nil
}

// CreateAllTables creates all table if not exists
func CreateAllTables(db *pg.DB) error {
	if err := CreateProductTable(db); err != nil {
		return err
	}

	return nil
}
