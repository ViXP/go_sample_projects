// Package adapter implements the Adapter design pattern.
// The space ship uses the default KeyPanel controller which conforms to a common Controller interface to receive the
// input from the pilot. A new joystick-style controller, SpaceStick, is introduced with its own incompatible interface.
// The StickToKeysAdapter adapts the SpaceStick to the legacy Controller interface, allowing the pilot to control the
// ship as if using the original controls.
package adapter

func Run() {
	ship := Ship{"Event Horizon", &KeyPanel{}}

	// Use old controller
	ship.Control.Up()
	ship.Control.Down()
	ship.Control.Left()
	ship.Control.Right()

	// Use the new controller through the adapter:
	ship.Control = &StickToKeysAdapter{joystick: &SpaceStick{}}
	ship.Control.Up()
	ship.Control.Down()
	ship.Control.Left()
	ship.Control.Right()
}
