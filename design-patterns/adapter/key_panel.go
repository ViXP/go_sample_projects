package adapter

import "fmt"

// KeyPanel is the standard main controller for the ship
type KeyPanel struct{}

// Up moves ship forward
func (kp *KeyPanel) Up() {
	fmt.Println("Will move forward")
}

// Down moves ship backward
func (kp *KeyPanel) Down() {
	fmt.Println("Will move backward")
}

// Left tilts ship to the left
func (kp *KeyPanel) Left() {
	fmt.Println("Will tilt left")
}

// Right tilts ship to the right
func (kp *KeyPanel) Right() {
	fmt.Println("Will tilt right")
}

// Explicit compile time check for the interface implementation
var _ Controller = &KeyPanel{}
