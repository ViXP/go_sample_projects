package composite

import (
	"fmt"
	"strings"
)

// Task represents the composite component.
type Task struct {
	WorkItem
	Children []DetailsPresenter
}

// GetStatus returns the human readable form of the component's status.
func (t *Task) GetStatus() (humanize string) {
	switch t.status {
	case todo:
		humanize = "To do"
	case in_progress:
		humanize = "In progress"
	case in_review:
		humanize = "In review"
	case done:
		humanize = "Done"
	default:
		humanize = "In backlog"
	}
	return
}

// NextStatus updates the status of the component to the next step in the pipeline.
func (t *Task) NextStatus() StatusManager {
	if t.status < done {
		t.status++
	}
	return t
}

// ShowDetails returns the aggregated details of the component and its sub-components.
func (t *Task) ShowDetails() string {
	var details strings.Builder
	details.WriteString(fmt.Sprintf("* Task '%s' - status: %s\n", t.Title, t.GetStatus()))

	for _, st := range t.Children {
		details.WriteString(st.ShowDetails())
	}
	return details.String()
}

// AddChild allows to add the children to the composite
func (t *Task) AddChild(child DetailsPresenter) Composer {
	t.Children = append(t.Children, child)
	return t
}

// Interface implementation assertion
var (
	_ StatusManager    = &Task{}
	_ DetailsPresenter = &Task{}
	_ Composer         = &Task{}
)
