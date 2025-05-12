package facade

import "fmt"

// MediaServer represents the devices that is able to broadcast the movie to the output interface
type MediaServer struct {
	name            string
	OutputInterface uint8
}

// On allows to turn on the media server
func (s *MediaServer) On() {
	fmt.Printf("Turning on the %s media server\n", s.name)
}

// Off allows to turn off the media server
func (s *MediaServer) Off() {
	fmt.Printf("Turning off the %s media server\n", s.name)
}

// Play allows to start playing the movie
func (s *MediaServer) Play(movieTitle string) {
	fmt.Printf("Playing the movie: %s\n", movieTitle)
}

// Stop allows to stop the movie
func (s *MediaServer) Stop() {
	fmt.Printf("Stop playing\n")
}

// Interfaces implementation assertion
var (
	_ Switcher = &MediaServer{}
	_ Player   = &MediaServer{}
)
