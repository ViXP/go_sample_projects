package data

import (
	"context"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int64     `json:"id"`
	Email    string    `json:"email"`
	Password string    `json:"-"`
	Name     string    `json:"first_name"`
	Surname  string    `json:"last_name"`
	Active   bool      `json:"active"`
	Created  time.Time `json:"created_at"`
	Updated  time.Time `json:"updated_at"`
}

func (u *User) All() ([]*User, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), timeout)

	defer cancelFunc()

	query := "SELECT id, email, password, first_name, last_name, active, created_at, updated_at FROM users ORDER BY id"

	rows, err := dbPool.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []*User

	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Email, &user.Password, &user.Name, &user.Surname, &user.Active, &user.Created, &user.Updated)

		if err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	return users, nil
}

func (u *User) Find(id int64) (*User, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), timeout)

	defer cancelFunc()

	query := "SELECT id, email, password, first_name, last_name, active, created_at, updated_at FROM users WHERE id = $1"

	row := dbPool.QueryRowContext(ctx, query, id)

	if row == nil {
		return nil, nil
	}

	var user User

	err := row.Scan(&user.ID, &user.Email, &user.Password, &user.Name, &user.Surname, &user.Active, &user.Created, &user.Updated)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *User) Create(user *User) (int, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), timeout)

	defer cancelFunc()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)

	if err != nil {
		return 0, err
	}

	query := "INSERT INTO users (email, password, first_name, last_name, active, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id"

	var id int
	err = dbPool.QueryRowContext(ctx, query, user.Email, hashedPassword, user.Name, user.Surname, user.Active, time.Now(), time.Now()).Scan(&id)

	return id, err
}

func (u *User) Update() error {
	ctx, cancelFunc := context.WithTimeout(context.Background(), timeout)

	defer cancelFunc()

	query := "UPDATE users SET email = $1, password = $2, first_name = $3, last_name = $4, active = $5, updated_at = $6 WHERE id = $7"

	_, err := dbPool.ExecContext(ctx, query, u.Email, u.Password, u.Name, u.Surname, u.Active, time.Now(), u.ID)

	return err
}

func (u *User) Destroy() error {
	ctx, cancelFunc := context.WithTimeout(context.Background(), timeout)

	defer cancelFunc()

	query := "DELETE FROM users WHERE id = $1"

	_, err := dbPool.ExecContext(ctx, query, u.ID)

	return err
}

func (u *User) FindByEmail(email string) (*User, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), timeout)

	defer cancelFunc()

	query := "SELECT id, email, password, first_name, last_name, active, created_at, updated_at FROM users WHERE email = $1"

	row := dbPool.QueryRowContext(ctx, query, email)

	if row == nil {
		return nil, nil
	}

	var user User

	err := row.Scan(&user.ID, &user.Email, &user.Password, &user.Name, &user.Surname, &user.Active, &user.Created, &user.Updated)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *User) ChangePassword(newPassword string) error {
	ctx, cancelFunc := context.WithTimeout(context.Background(), timeout)

	defer cancelFunc()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), 12)

	if err != nil {
		return err
	}

	query := "UPDATE users SET password = $1, updated_at = $2 WHERE id = $3"

	_, err = dbPool.ExecContext(ctx, query, hashedPassword, time.Now(), u.ID)

	return err
}

func (u *User) IsCorrectPassword(comparedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(comparedPassword))
	return err == nil
}
