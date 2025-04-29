package adapter

import "fmt"

// Direction enum
const (
	forward uint = iota
	backward
	tilt_left
	tilt_right
)

// SpaceStick is the new modern controller with the custom interface
type SpaceStick struct{}

// Touch moves the object in the corresponding direction
func (s *SpaceStick) Touch(direction uint) {
	var move string
	switch direction {
	case forward:
		move = "move forward"
	case backward:
		move = "move backward"
	case tilt_left:
		move = "tilt left"
	case tilt_right:
		move = "tilt right"
	default:
		move = "paused"
	}

	fmt.Printf("Will %s\n", move)
}
