package main

import (
	"fmt"

	"example.com/structs-demo/user"
)

func getUserInput(message string) (input string) {
	fmt.Printf("%v: ", message)
	fmt.Scanln(&input)
	return
}

func main() {
	firstName := getUserInput("Enter your first name")
	lastName := getUserInput("Enter your last name")
	birthDate := getUserInput("Enter your birthdate")

	customUser, err := user.New(firstName, lastName, birthDate)

	if err != nil {
		panic(err)
	}

	customAdmin, err := user.NewAdmin("sample@mail.com", "12345a")

	if err != nil {
		panic(err)
	}

	customUser.OutputData()
	customUser.ClearName()
	customAdmin.OutputData()
}
