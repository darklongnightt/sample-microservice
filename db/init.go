package db

import (
	"fmt"
	"time"

	"github.com/go-pg/pg"
)

// Init ...
func Init() (*pg.DB, error) {
	var db *pg.DB
	connectionTries := 5

	for connectionTries > 0 {
		// Setup db connection
		db = pg.Connect(&pg.Options{
			Addr:     "db:5432",
			User:     "postgres",
			Password: "password",
			Database: "sample",
		})

		// Check if connection is ok
		if _, err := db.Exec("SELECT 1"); err != nil {
			connectionTries--
			fmt.Printf("failed to connect to db, tries left: %v\n", connectionTries)
			time.Sleep(5 * time.Second)

			if connectionTries == 0 {
				return nil, fmt.Errorf("failed to connect to db\nreason: %v", err)
			}
		} else {
			break
		}
	}

	// Init all tables
	fmt.Println("Connection to db successful")
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
