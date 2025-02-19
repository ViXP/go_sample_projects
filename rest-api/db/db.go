package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("no connection to the database")
	}

	DB.SetMaxOpenConns(20)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createUsers()
	createEvents()
}

func createEvents() {
	createEvents := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name VARCHAR NOT NULL,
		description TEXT NOT NULL,
		location VARCHAR NOT NULL,
		date_time DATETIME NOT NULL,
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`
	_, err := DB.Exec(createEvents)

	if err != nil {
		panic("can't create table events")
	}
}

func createUsers() {
	createUsers := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email VARCHAR NOT NULL UNIQUE,
		password VARCHAR NOT NULL
	)
	`
	_, err := DB.Exec(createUsers)

	if err != nil {
		panic("can't create table users")
	}
}
