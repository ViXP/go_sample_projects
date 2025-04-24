// Package prototype implements the Prototype design pattern.
// NewCharacter represents the prototyped constructor of NPC in the abstract game, that creates the character
// based on the general behavior of the specified character and assigns additional optional properties to
// the new character.
package prototype

import "fmt"

func Run() {
	// Creating the friendly NPC from prototype
	farmer := NewCharacter(Friendly, "Farmer", -1, nil)

	// Creating the foe NPC from prototype with custom shape
	bandit := NewCharacter(Hostile, "Bandit", 10, &Shape{
		Head: ShapePart{
			Color: "#f00",
			model: Cone,
			Size:  1.5,
			Count: 1},
		Body: ShapePart{
			Color: "#eee",
			model: Cylinder,
			Size:  2,
			Count: 1,
		},
		Legs: ShapePart{
			Color: "#440",
			model: Cylinder,
			Size:  2.5,
			Count: 2,
		},
	})

	// Output the info (using Stringer interface)
	fmt.Println(farmer)
	fmt.Println(bandit)
}
