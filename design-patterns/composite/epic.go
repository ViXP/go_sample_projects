package composite

import (
	"fmt"
	"strings"
)

// Epic represents the composite component.
type Epic struct {
	WorkItem
	Children []DetailsPresenter
}

// GetStatus returns the human readable form of the component's status.
func (e *Epic) GetStatus() (humanize string) {
	switch e.status {
	case in_progress:
		humanize = "In progress"
	case done:
		humanize = "Done"
	default:
		humanize = "To do"
	}
	return
}

// NextStatus updates the status of the component to the next step in the pipeline.
func (e *Epic) NextStatus() StatusManager {
	if e.status < done {
		e.status++
	}
	return e
}

// ShowDetails returns the aggregated details of the component and its sub-components.
func (e *Epic) ShowDetails() string {
	var details strings.Builder
	details.WriteString(fmt.Sprintf("Epic '%s' - status: %s\n", e.Title, e.GetStatus()))

	for _, t := range e.Children {
		details.WriteString(t.ShowDetails())
	}
	return details.String()
}

// AddChild allows to add the children to the composite
func (e *Epic) AddChild(child DetailsPresenter) Composer {
	e.Children = append(e.Children, child)
	return e
}

// Interface implementation assertion
var (
	_ StatusManager    = &Epic{}
	_ DetailsPresenter = &Epic{}
	_ Composer         = &Epic{}
)
