package prototype

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

// Behavioral constants enum
const (
	Hostile uint16 = iota
	Friendly
)

// Prototyper is the interface that is implemented by Character to create the cloned struct based on the existing one
type Prototyper interface {
	DeepClone() (Prototyper, error)
}

// Character is a complicated composed structure that represents a NPC in the game
type Character struct {
	Name     string
	behavior uint16
	Speed    float32
	Shape    *Shape
}

// DeepClone implements Prototyper interface and allows to create a deep copy of character through binary serialization
func (ch *Character) DeepClone() (Prototyper, error) {
	var cloned Character
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	decoder := gob.NewDecoder(&buffer)
	err := encoder.Encode(&ch)

	if err != nil {
		return nil, err
	}

	err = decoder.Decode(&cloned)

	if err != nil {
		return nil, err
	}

	return &cloned, nil
}

// Behavior is the human readable getter of the behavior property of the character
func (ch *Character) Behavior() string {
	switch ch.behavior {
	case Hostile:
		return "Enemy"
	case Friendly:
		return "Friend"
	default:
		return "unknown"
	}
}

func (ch *Character) String() string {
	return fmt.Sprintf("I am %s\nMy speed is %v km/h and I am your %s\nI look like:\n  %v\n", ch.Name, ch.Speed, ch.Behavior(), ch.Shape)
}

// Explicit compile time checks if the pointer to the Character struct implements Prototyper and Stringer interfaces
var _ Prototyper = (*Character)(nil)
var _ fmt.Stringer = (*Character)(nil)

// NewCharacter is the Character constructor based on the prototypes selected by behavior
func NewCharacter(behavior uint16, name string, speed float32, shape *Shape) Prototyper {
	var npc Prototyper
	var err error

	if behavior == Hostile {
		npc, err = EnemyPrototype.DeepClone()
	} else {
		npc, err = FriendPrototype.DeepClone()
	}

	if err != nil {
		panic(err)
	}

	if char, ok := npc.(*Character); ok {
		char.Name = name

		if speed >= 0 {
			char.Speed = speed
		}

		if shape != nil {
			char.Shape = shape
		}
		return char
	}

	return npc
}
