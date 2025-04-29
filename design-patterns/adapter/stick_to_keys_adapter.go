package adapter

// StickToKeysAdapter is the interface, that implements Controller interface for the new type of joystick controllers
type StickToKeysAdapter struct {
	joystick *SpaceStick
}

// Up moves ship forward
func (a *StickToKeysAdapter) Up() {
	a.joystick.Touch(forward)
}

// Down moves ship backward
func (a *StickToKeysAdapter) Down() {
	a.joystick.Touch(backward)
}

// Left tilts ship to the left
func (a *StickToKeysAdapter) Left() {
	a.joystick.Touch(tilt_left)
}

// Right tilts ship to the right
func (a *StickToKeysAdapter) Right() {
	a.joystick.Touch(tilt_right)
}

// Explicit check for the interface implementation
var _ Controller = &StickToKeysAdapter{}
