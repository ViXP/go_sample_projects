package factory

import "fmt"

// The limits of the charge for the fundamental particles
const (
	MaxCharge = 1
	MinCharge = -1
)

// Particles classes enum constants
const (
	Boson uint16 = iota
	Fermion
	Composite
)

// Particle is the general struct that represents particle model
type Particle struct {
	class  uint16
	Spin   float32
	Mass   float32
	Charge int16
}

// AddMass is the general interface for the mass reassignment of the particle
func (p *Particle) AddMass(mass float32) *Particle {
	p.Mass = mass
	return p
}

// AddCharge is the general interface for the charge assignment of the particle (with validations)
func (p *Particle) AddCharge(charge int16) (*Particle, error) {
	if p.class != Composite && (charge > MaxCharge || charge < MinCharge) {
		return nil, fmt.Errorf("impossible charge")
	}
	p.Charge = charge
	return p, nil
}

// Class returns the human readable type of the Particle
func (p *Particle) Class() string {
	switch p.class {
	case Boson:
		return "Boson"
	case Fermion:
		return "Fermion"
	case Composite:
		return "Composite (not specified)"
	default:
		return "Unknown!"
	}
}

// String is a stringified representation of the Particle including it's known attributes
func (p *Particle) String() string {
	return fmt.Sprintf("Particle type: %s\nSpin: %v\nMass: %v MeV/c^2\nCharge: %v\n", p.Class(), p.Spin, p.Mass, p.Charge)
}
