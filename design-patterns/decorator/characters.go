package decorator

// NewKnight is a factory function that creates a new knight.
func NewKnight() *Character {
	return &Character{Class: "Knight", Speed: 8.3, Attack: 30.0, Defense: 60}
}

// NewMagician is a factory function that creates a new magician.
func NewMagician() *Character {
	return &Character{Class: "Magician", Speed: 10.2, Attack: 20.5, Defense: 80}
}

// NewHorseman is a factory function that creates a new horseman.
func NewHorseman() *Character {
	return &Character{Class: "Horseman", Speed: 25.2, Attack: 10.5, Defense: 30}
}
