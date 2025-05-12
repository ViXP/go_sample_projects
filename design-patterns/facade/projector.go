package facade

import "fmt"

// Possible inputs enum
const (
	hdmi uint8 = iota
	displayPort
	vga
	rgb
	sVideo
	scart
	rca
)

// Projector represents the projection device of the home theater
type Projector struct {
	name         string
	currentInput uint8
}

// On allows to turn on the projector
func (p *Projector) On() {
	fmt.Printf("Turning on the %s\n", p.name)
}

// Off allows to turn off the projector
func (p *Projector) Off() {
	fmt.Printf("Turning off the %s\n", p.name)
}

// SetInput allows to switch the active input of the projector
func (p *Projector) SetInput(input uint8) {
	var parsedInput string
	switch input {
	case hdmi:
		parsedInput = "HDMI"
	case displayPort:
		parsedInput = "Display Port"
	case vga:
		parsedInput = "VGA"
	case rgb:
		parsedInput = "RGB"
	case sVideo:
		parsedInput = "SVideo"
	case scart:
		parsedInput = "SCART"
	case rca:
		parsedInput = "RCA"
	default:
		return
	}

	p.currentInput = input

	fmt.Printf("Switching the %s to the %s input\n", p.name, parsedInput)
}

// Interface implementation assertion
var _ Switcher = &Projector{}
