package factory

import "math"

// ParticleFactory is the abstract factory interface
type ParticleFactory interface {
	Create(float32, float32, int16) (*Particle, error)
}

// FundamentalParticleFactory is the specific factory implementation
type FundamentalParticleFactory struct{}

func (f *FundamentalParticleFactory) Create(spin float32, mass float32, charge int16) (*Particle, error) {
	var particle Particle

	if math.Mod(float64(spin), 1) == 0 {
		particle = Particle{
			class: Boson,
			Spin:  spin,
		}
	} else {
		particle = Particle{
			class: Fermion,
			Spin:  spin,
		}
	}

	particle.AddMass(mass)
	p, err := particle.AddCharge(charge)

	if err != nil {
		return nil, err
	}

	return p, nil
}

// CompositeParticleFactory is a simple implementation of a specific factory to create the composite particle.
// There is no distinguish of a specific particle types here (mesons, baryons, hadrons etc.)
type CompositeParticleFactory struct{}

func (f *CompositeParticleFactory) Create(spin float32, mass float32, charge int16) (*Particle, error) {
	particle := Particle{
		class: Composite,
		Spin:  spin,
		Mass:  mass,
	}
	p, err := particle.AddCharge(charge)

	if err != nil {
		return nil, err
	}

	return p, nil
}
