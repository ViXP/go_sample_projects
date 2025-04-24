package prototype

import "fmt"

// 3D figures enum that represent the model used for character's shape
const (
	Cube uint16 = iota
	Sphere
	Cylinder
	Cone
	Pyramid
)

// ShapePart is the struct that encapsulates the data required to build the parts of the character's shape
type ShapePart struct {
	Size  float32
	model uint16
	Color string
	Count uint16
}

// Model is the human readable getter of the model property of the shape part
func (sp *ShapePart) Model() string {
	switch sp.model {
	case Cube:
		return "cubic"
	case Sphere:
		return "spherical"
	case Cylinder:
		return "cylindrical"
	case Cone:
		return "conical"
	case Pyramid:
		return "pyramidal"
	default:
		return "unknown"
	}
}

func (sp *ShapePart) String() string {
	return fmt.Sprintf("Size: %v; Shape: %s; Color: %v; Quantity: %v", sp.Size, sp.Model(), sp.Color, sp.Count)
}
