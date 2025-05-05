package composite

// Status values for a WorkItem
const (
	backlog uint8 = iota
	todo
	in_progress
	in_review
	done
)

// StatusManager defines behavior for retrieving and updating the workflow status of a component.
type StatusManager interface {
	GetStatus() string
	NextStatus() StatusManager
}

// DetailsPresenter defines behavior for displaying component details, including nested ones.
type DetailsPresenter interface {
	ShowDetails() string
}

// Composer is implemented by composite components only that can contain child components.
type Composer interface {
	AddChild(DetailsPresenter) Composer
}

// WorkItem holds the common fields shared by both leaf and composite components.
type WorkItem struct {
	Title  string
	status uint8
}
