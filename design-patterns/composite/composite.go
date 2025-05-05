// Package composite implements the Composite design pattern.
// It models classic Agile concepts such as Epics, Tasks, and SubTasks.
// SubTasks are leaf nodes, while Epics and Tasks are composites that can contain other components—either leaf nodes or
// nested composites. All components implement a shared set of interfaces.
package composite

import "fmt"

func Run() {
	epic := Epic{WorkItem: WorkItem{Title: "My first epic"}}
	task := Task{WorkItem: WorkItem{Title: "First task in epic"}}
	subtask1 := SubTask{WorkItem: WorkItem{Title: "Do some work"}}
	subtask2 := SubTask{WorkItem: WorkItem{Title: "Make some other work"}}
	epic.AddChild(&task)
	task.AddChild(&subtask1).AddChild(&subtask2)

	epic.NextStatus().NextStatus()
	task.NextStatus().NextStatus()
	subtask2.NextStatus().NextStatus()

	fmt.Println(epic.ShowDetails())
}
