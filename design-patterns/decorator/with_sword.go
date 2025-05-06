package decorator

import "fmt"

// WithSword is the character's decorator for enhanced attack.
type WithSword struct {
	Actor
}

// Hit executes the enhanced attack
func (d *WithSword) Hit() string {
	return fmt.Sprintf("%s + Additional damage from sword +%v!", d.Actor.Hit(), d.Effectiveness())
}

// Effectiveness returns the effectiveness number of the Buff.
func (d *WithSword) Effectiveness() uint8 {
	return 10
}

// Interfaces implementation assertion.
var (
	_ Actor = &WithSword{}
	_ Buff  = &WithSword{}
)

// EquipSword is the factory function for the WithSword decorator.
func EquipSword(a Actor) *WithSword {
	return &WithSword{a}
}
