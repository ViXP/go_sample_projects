package models

import (
	"errors"

	"example.com/rest-api/db"
)

type Registration struct {
	ID      int64
	EventID int64
	UserID  int64
}

func GetRegistrationByEventUserId(eventId, userId int64) (*Registration, error) {
	query := "SELECT * FROM registrations WHERE event_id = ? AND user_id = ? LIMIT 1"
	row := db.DB.QueryRow(query, eventId, userId)

	var registration Registration

	err := row.Scan(&registration.ID, &registration.EventID, &registration.UserID)

	if err != nil {
		return nil, errors.New("fetch error")
	}

	return &registration, nil
}

func (r *Registration) Save() error {
	query := `
		INSERT INTO registrations(user_id, event_id)
		VALUES (?, ?)
	`
	prepared, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer prepared.Close()

	result, err := prepared.Exec(r.UserID, r.EventID)

	if err != nil {
		return err
	}

	r.ID, err = result.LastInsertId()
	return err
}

func (r *Registration) Destroy() error {
	query := `
		DELETE FROM registrations
		WHERE id = ?
	`
	prepared, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer prepared.Close()

	_, err = prepared.Exec(r.ID)

	return err
}
