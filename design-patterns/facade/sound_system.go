package facade

import "fmt"

// SoundSystem represents the sound system of the home theater
type SoundSystem struct {
	name   string
	volume uint8
}

// On allows to turn on the sound system
func (s *SoundSystem) On() {
	fmt.Printf("Turning on the %s sound system\n", s.name)
}

// Off allows to turn off the sound system
func (s *SoundSystem) Off() {
	fmt.Printf("Turning off the %s sound system\n", s.name)
}

// SetVolume allows to set the specific volume on the speakers
func (s *SoundSystem) SetVolume(volume uint8) {
	s.volume = volume
	fmt.Printf("Set volume to: %v%%\n", s.volume)
}

// Interface implementation assertion
var _ Switcher = &SoundSystem{}
