package adapter

// Controller is the common interface expected for the ship's controller
type Controller interface {
	Up()
	Down()
	Left()
	Right()
}
