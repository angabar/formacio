package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error

	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Error connecting to database")
	}

	// Conexiones maximas abiertas al mismo tiempo
	DB.SetMaxOpenConns(10)

	// Conexiones maximas abiertas y sin uso al mismo tiempo
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createEventsTable := `
		CREATE TABLE IF NOT EXISTS events (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			description TEXT NOT NULL,
			location TEXT NOT NULL,
			dateTime DATETIME NOT NULL,
			user_id INTEGER
		)
	`

	_, err := DB.Exec(createEventsTable)

	if err != nil {
		fmt.Println(err)
		panic("Error creating events table")
	}
}
