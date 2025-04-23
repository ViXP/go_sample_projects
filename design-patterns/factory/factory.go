// Package factory implements the Factory and Abstract Factory design pattern.
// The interface ParticleFactory declares the abstract factory interface, while FundamentalParticleFactory and
// CompositeParticleFactory are the specific implementations of this abstract factory, which allows us to
// produce a simple models that represent untitled fundamental and composite particles of physical world.
package factory

import (
	"fmt"
)

func Run() {
	fundamentalFactory := FundamentalParticleFactory{}
	compositeFactory := CompositeParticleFactory{}

	// Creating the Electron through fundamental factory:
	electron, err := fundamentalFactory.Create(0.5, 0.511, -1)

	if err != nil {
		panic(err)
	}

	// Creating the Higgs Boson through fundamental factory:
	higgsp, err := fundamentalFactory.Create(0, 125_000, 0)

	if err != nil {
		panic(err)
	}

	// Creating the Proton through composite factory
	proton, err := compositeFactory.Create(0.5, 938.272, 1)

	if err != nil {
		panic(err)
	}

	// Let's output the info:
	fmt.Println(electron)
	fmt.Println(higgsp)
	fmt.Println(proton)
}
