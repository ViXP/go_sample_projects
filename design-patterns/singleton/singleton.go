// Package singleton implements the Singleton design pattern.
// The example demonstrates the usability of the Notepad struct, that can be used in different places in the application
// through the GetTodoNotepad singleton creating method.
package singleton

import "fmt"

func Run() {
	n := GetTodoNotepad()

	n.AddNote("Some action").AddNote("Some other action")

	// Will work with the same notepad but in different function, without dependency injection
	addOtherActions()

	fmt.Println(n)
}

func addOtherActions() {
	n := GetTodoNotepad()
	n.AddNote("Completely different action, added from elsewhere")
}
