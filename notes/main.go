package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/notes/note"
	"example.com/notes/todo"
)

type saver interface {
	Save() error
}

type outputtable interface {
	saver
	Display()
}

func main() {
	for {
		var answer string
		title, _ := getUserInput("Note title")
		content, _ := getUserInput("Note content")

		userNote, err := note.New(title, content)

		if err != nil {
			panic(err)
		}

		outputData(userNote)

		fmt.Print("Do you wish to continue? [y/n] ")
		fmt.Scanln(&answer)

		if strings.ToLower(answer) == "n" {
			break
		}
	}

	todoText, _ := getUserInput("Add something to todo")
	userTodo, err := todo.New(todoText)

	if err != nil {
		panic(err)
	}

	outputData(userTodo)
}

func saveData(data saver) {
	if data.Save() != nil {
		panic("The todo could not be saved to file!")
	}
}

func outputData(data outputtable) {
	data.Display()
	saveData(data)
}

func getUserInput(prompt string) (input string, error error) {
	fmt.Printf("%v: ", prompt)

	reader := bufio.NewReader(os.Stdin)

	input, error = reader.ReadString('\n')

	input = strings.TrimSuffix(input, "\n")

	return
}
