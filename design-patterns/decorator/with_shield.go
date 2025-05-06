package decorator

import "fmt"

// WithShield is the character's decorator for enhanced defense.
type WithShield struct {
	Actor
}

// Block executes the enhanced block.
func (d *WithShield) Block() string {
	return fmt.Sprintf("%s + Additional shield protection +%v%%!", d.Actor.Block(), d.Effectiveness())
}

// Effectiveness returns the effectiveness number of the Buff.
func (d *WithShield) Effectiveness() uint8 {
	return 15
}

// Interfaces implementation assertion.
var (
	_ Actor = &WithShield{}
	_ Buff  = &WithShield{}
)

// EquipShield is the factory function for the WithShield decorator.
func EquipShield(a Actor) *WithShield {
	return &WithShield{a}
}
