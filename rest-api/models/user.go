package models

import (
	"errors"

	"example.com/rest-api/db"
	"example.com/rest-api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := `
		INSERT INTO users(email, password)
		VALUES (?, ?)
	`
	prepared, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer prepared.Close()

	hashedPass, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	result, err := prepared.Exec(u.Email, hashedPass)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	u.ID = id

	return err
}

func (u *User) VerifyPassword() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var dbPassword string

	err := row.Scan(&u.ID, &dbPassword)

	if err != nil {
		return errors.New("user is not found")
	}

	if utils.ComparePasswordHash(u.Password, dbPassword) {
		return nil
	} else {
		return errors.New("wrong password")
	}
}
