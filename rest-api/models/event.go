package models

import (
	"errors"
	"time"

	"example.com/rest-api/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      any
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, errors.New("can't retrieve the records")
	}

	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

		if err != nil {
			return nil, errors.New("can't parse fetched records")
		}

		events = append(events, event)
	}

	return events, nil
}

func GetEventByID(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ? LIMIT 1"

	row := db.DB.QueryRow(query, id)

	var event Event

	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

	if err != nil {
		return nil, errors.New("fetch error")
	}

	return &event, nil
}

func (e *Event) Save() error {
	query := `
		INSERT INTO events(name, description, location, date_time)
		VALUES (?, ?, ?, ?)
	`
	prepared, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer prepared.Close()

	result, err := prepared.Exec(e.Name, e.Description, e.Location, e.DateTime)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	e.ID = id

	return err
}

func (e *Event) Update() error {
	query := `
		UPDATE events
		SET name = ?, description = ?, location = ?, date_time = ?
		WHERE id = ?
	`

	prepared, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer prepared.Close()

	_, err = prepared.Exec(e.Name, e.Description, e.Location, e.DateTime, e.ID)

	return err
}

func (e *Event) Destroy() error {
	query := `
		DELETE FROM events
		WHERE id = ?
	`

	prepared, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer prepared.Close()

	_, err = prepared.Exec(e.ID)

	return err
}
