package prototype

import "fmt"

// Shape is the struct that represents the shape of the Character
type Shape struct {
	Head ShapePart
	Body ShapePart
	Legs ShapePart
}

func (s *Shape) String() string {
	return fmt.Sprintf("Head: %v;\n  Body: %v;\n  Legs: %v;", s.Head.String(), s.Body.String(), s.Legs.String())
}
