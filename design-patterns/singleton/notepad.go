package singleton

import (
	"bytes"
	"fmt"
	"sync"
)

// Notepad represents the object, that encapsulates the logic of adding and printing the notes into it
type Notepad struct {
	title string
	notes []string
}

func (n *Notepad) printNotes() string {
	bb := new(bytes.Buffer)
	for i, note := range n.notes {
		bb.WriteString(fmt.Sprintf("%v. %s\n", i+1, note))
	}

	return bb.String()
}

// String implement Stringer interface for the notepad to be able to work with the elements in print ready format
func (n *Notepad) String() string {
	return fmt.Sprintf("%s:\n%s", n.title, n.printNotes())
}

// AddNote is the part of the public interface for the notes mutations
func (n *Notepad) AddNote(note string) *Notepad {
	n.notes = append(n.notes, note)
	return n
}

var todoNotepad *Notepad
var once sync.Once

// GetTodoNotepad is the Singleton constructor for the Todo list notepad, that can be reused in different places of
// the application
func GetTodoNotepad() *Notepad {
	once.Do(func() {
		if todoNotepad == nil {
			todoNotepad = &Notepad{"Todo list", nil}
		}
	})
	return todoNotepad
}
