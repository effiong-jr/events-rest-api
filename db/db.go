package db

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() {

	var err error

	DB, err = sql.Open("sqlite", "api.db")

	if err != nil {
		panic("Database connection failed!")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {

	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT UNIQUE NOT NULL,
		password TEXT NOT NULL
	)
	`

	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        description TEXT NOT NULL, 
        location TEXT NOT NULL,
        dateTime DATETIME NOT NULL,
        userId INTEGER,
		FOREIGN KEY(userId) REFERENCES users(id)
    )
	`

	_, err := DB.Exec(createUsersTable)
	if err != nil {
		panic("Could not create users table")
	}

	_, err = DB.Exec(createEventsTable)

	if err != nil {
		panic("Could not create events table")
	}

}