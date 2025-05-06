package decorator

import "fmt"

// Character is the common struct of all RPG characters.
type Character struct {
	Class   string
	Speed   float32
	Attack  float32
	Defense uint8
}

// Hit executes the attack.
func (c *Character) Hit() string {
	return fmt.Sprintf("%s attacks with %f power!", c.Class, c.Attack)
}

// Block defends the character from the attack.
func (c *Character) Block() string {
	return fmt.Sprintf("%s is able to block the attack with %v%% effectiveness!", c.Class, c.Defense)
}

// Walk allows the character to move toward the battlefield.
func (c *Character) Walk() string {
	return fmt.Sprintf("%s walks with the speed of %v km/h!", c.Class, c.Speed)
}

// Interface implementation assertion.
var _ Actor = &Character{}
