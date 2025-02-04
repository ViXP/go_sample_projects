package userinput

import (
	"encoding/json"
	"fmt"
	"os"
)

type UserInput struct{}

func New() *UserInput {
	return &UserInput{}
}

func (*UserInput) ReadLines() ([]string, error) {
	var unparsedStrings []string

	fmt.Println("Please, enter the values (0 for exit):")

	for {
		var inputted string
		fmt.Scan(&inputted)

		if inputted == "0" {
			break
		}
		unparsedStrings = append(unparsedStrings, inputted)
	}

	return unparsedStrings, nil
}

func (*UserInput) WriteJson(data interface{}) error {
	encoder := json.NewEncoder(os.Stdout)
	err := encoder.Encode(data)

	if err != nil {
		return fmt.Errorf("data is not JSON-able")
	}
	return nil
}
