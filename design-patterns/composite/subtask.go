package composite

import "fmt"

// SubTask represents the leaf component.
type SubTask struct {
	WorkItem
}

// GetStatus returns the human readable form of the component's status.
func (st *SubTask) GetStatus() (humanize string) {
	switch st.status {
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
func (st *SubTask) NextStatus() StatusManager {
	if st.status < done {
		st.status++
	}

	return st
}

// ShowDetails returns the details of the component.
func (st *SubTask) ShowDetails() string {
	return fmt.Sprintf("** Sub Task '%s' - status: %s\n", st.Title, st.GetStatus())
}

// Interface implementation assertion
var (
	_ StatusManager    = &SubTask{}
	_ DetailsPresenter = &SubTask{}
)
