package decorator

import "fmt"

// WithPotion is the character's decorator for enhanced speed.
type WithPotion struct {
	Actor
}

// Walk executes the enhanced walk.
func (d *WithPotion) Walk() string {
	return fmt.Sprintf("%s + Additional speed increase from Speed potion +%v km/h!", d.Actor.Walk(), d.Effectiveness())
}

// Effectiveness returns the effectiveness number of the Buff.
func (d *WithPotion) Effectiveness() uint8 {
	return 25
}

// Interfaces implementation assertion.
var (
	_ Actor = &WithPotion{}
	_ Buff  = &WithPotion{}
)

// UsePotion is the factory function for the WithPotion decorator.
func UsePotion(a Actor) *WithPotion {
	return &WithPotion{a}
}
