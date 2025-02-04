package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func (u *User) OutputData() {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func (u *User) ClearName() {
	u.firstName = ""
	u.lastName = ""
}

type Admin struct {
	email    string
	password string
	User
}

func New(firstName, lastName, birthDate string) (*User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("first name, last name and birthdate are required")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}

func NewAdmin(email, password string) (*Admin, error) {
	if email == "" || password == "" {
		return nil, errors.New("email and password are required")
	}

	return &Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "ADMIN",
			lastName:  "ADMIN",
			birthDate: "---",
			createdAt: time.Now(),
		},
	}, nil
}
