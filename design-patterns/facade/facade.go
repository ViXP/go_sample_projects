// Package facade implements the Facade design pattern.
// HomeTheater provides a simplified interface for controlling a group of devices: Projector, MediaServer, Lights,
// and SoundSystem. It encapsulates the logic required to coordinate these components in the correct sequence,
// allowing viewers to prepare for, play, and shut down a movie session through a single high-level interface.
package facade

import "time"

func Run() {
	homeTheater := NewHomeTheater(displayPort, "Generic 4K Projector", "PC", "Generic LED Strip Lighting", "Dolby Atmos")

	homeTheater.Prepare()
	homeTheater.PlayMovie("Inception")

	time.Sleep(5 * time.Second)

	homeTheater.EndMovie()
	homeTheater.Close()
}
