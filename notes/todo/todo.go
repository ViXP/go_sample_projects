package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func (todo Todo) Display() {
	fmt.Printf("Your todo has the following content:\n\n%v\n", todo.Text)
}

func (todo Todo) Save() error {
	content, error := json.Marshal(todo)

	if error != nil {
		return error
	}

	return os.WriteFile("todo.json", content, 0644)
}

func New(content string) (*Todo, error) {
	if content == "" {
		return &Todo{}, errors.New("text is empty")
	}

	return &Todo{
		Text: content,
	}, nil
}
