// Package decorator implements the Decorator design pattern.
// In the role-playing game example, we have three character classes: Knight, Magician, and Horseman.
// Each character has predefined attributes represented by the Character struct.
// Decorators such as WithPotion, WithShield, and WithSword enhance a character's abilities.
// These decorators can be applied in a composable manner, allowing for dynamic stacking of enhancements
// without modifying the original character implementation.
package decorator

import "fmt"

func Run() {
	// Creating the unique characters with the factory functions (constructors)
	armoredKnight := EquipSword(EquipShield(NewKnight()))
	bladedMagician := EquipSword(NewMagician())
	enchantedHorseman := UsePotion(NewHorseman())

	// Actions execution
	fmt.Println(armoredKnight.Walk())
	fmt.Println(armoredKnight.Hit())
	fmt.Println(armoredKnight.Block())

	fmt.Println(bladedMagician.Walk())
	fmt.Println(bladedMagician.Hit())
	fmt.Println(bladedMagician.Block())

	fmt.Println(enchantedHorseman.Walk())
	fmt.Println(enchantedHorseman.Hit())
	fmt.Println(enchantedHorseman.Block())
}
