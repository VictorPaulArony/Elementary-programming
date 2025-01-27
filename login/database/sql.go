package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// function to initialize the db for table creations
func InitializeDb(dbFIlePath string) (*sql.DB, error) {
	// opening the db file
	db, err := sql.Open("sqlite3", dbFIlePath)
	if err != nil {
		return nil, err
	}
	// Creating the db for the system
	query := `
	CREATE TABLE IF NOT EXISTS users(
		id TEXT PRIMARY KEY,
		useremail TEXT UNIQUE NOT NULL,
		username TEXT UNIQUE NOT NULL,
		password TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`
	if _, err := db.Exec(query); err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}
