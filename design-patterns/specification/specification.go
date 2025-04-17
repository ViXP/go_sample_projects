// Package specification implements Specification tactical pattern.
// Validations are implemented as separate structs, each sharing a common interface that defines a single function.
// This allows us to utilize multiple, reusable validation rules independently.
// The `validUsers()` function remains unchanged while supporting different validation logic,
// leveraging polymorphism via injected specification dependencies.
// This approach adheres to the Open/Closed Principle — allowing behavior extension without modifying the existing code.
package specification

import (
	"fmt"
	"regexp"
)

type UserEntry struct {
	name  string
	email string
	phone string
}

func (u *UserEntry) String() string {
	return fmt.Sprintf("User name: %v, email: %v, phone: %v", u.name, u.email, u.phone)
}

// Validation is the common Specification interface
type Validation interface {
	IsValid(entry *UserEntry) bool
}

// PhoneValidation is a specific Specification implementation
type PhoneValidation struct{}

func (pv PhoneValidation) IsValid(entry *UserEntry) bool {
	return len(entry.phone) > 7
}

type EmailValidation struct{}

func (ev EmailValidation) IsValid(entry *UserEntry) bool {
	result, _ := regexp.MatchString(`^(\w|\d)+@(\w|\d)+\.(\w|\d)+$`, entry.email)
	return result
}

type NameValidation struct{}

func (nv NameValidation) IsValid(entry *UserEntry) bool {
	return len(entry.name) > 0
}

func validUsers(users *[]UserEntry, validations *[]Validation) []*UserEntry {
	validUsers := make([]*UserEntry, 0)
	for _, u := range *users {
		for _, v := range *validations {
			if !(v.IsValid(&u)) {
				goto next
			}
		}
		validUsers = append(validUsers, &u)
	next:
	}

	return validUsers
}

func Run() {
	users := []UserEntry{
		{"", "valid@mail.com", "555-12-14"},
		{"Valid User", "valid@too.com", "095959595"},
		{"Some Invalid User", "valid@mail.com", "invalid"},
	}

	valid := validUsers(&users, &[]Validation{PhoneValidation{}, EmailValidation{}, NameValidation{}})

	fmt.Println("SPECIFICATION")
	fmt.Println(valid)
}
