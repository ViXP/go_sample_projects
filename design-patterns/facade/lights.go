package facade

import "fmt"

// Lights describes the lighting equipment of the cinema
type Lights struct {
	description string
	brightness  uint8
}

// On allows to turn on the lights
func (l *Lights) On() {
	fmt.Printf("Turning on the %s\n", l.description)
}

// Off allows to turn off the lights
func (l *Lights) Off() {
	fmt.Printf("Turning off the %s\n", l.description)
}

// SetBrightness allows to dim or brighten up the lights to the specific absolute percentage
func (l *Lights) SetBrightness(amount uint8) {
	l.brightness = amount
	fmt.Printf("Dimming the %s to %v%%\n", l.description, l.brightness)
}

// Interface implementation assertion
var _ Switcher = &Lights{}
