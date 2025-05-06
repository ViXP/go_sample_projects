package decorator

// Actor defines the interface for all of the characters' actions.
type Actor interface {
	Block() string
	Hit() string
	Walk() string
}

// Buff defines the interface for all of the characters' decorators.
type Buff interface {
	Effectiveness() uint8
}
