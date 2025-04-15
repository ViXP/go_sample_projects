package three

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type triangle struct {
	a, b, c float64
}

type rectangle struct {
	h, w float64
}

type figure interface {
	area() float64
	fmt.Stringer
}

func (c circle) String() string {
	return fmt.Sprintf("Circle with radius: %.2f", c.radius)
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (t triangle) String() string {
	return fmt.Sprintf("Triangle with sides lengths: %.2f, %.2f, %.2f", t.a, t.b, t.c)
}

func (t triangle) area() float64 {
	// Heron's formula
	hp := (t.a + t.b + t.c) / 2.0
	return math.Sqrt(hp * (hp - t.a) * (hp - t.b) * (hp - t.c))
}

func (t triangle) angle(a, b, c float64) float64 {
	return math.Acos((a*a+b*b+-c*c)/(2*a*b)) * 180.0 / math.Pi
}

func (t triangle) angles() []float64 {
	return []float64{t.angle(t.a, t.b, t.c), t.angle(t.b, t.c, t.a), t.angle(t.c, t.a, t.b)}
}

func (r rectangle) String() string {
	return fmt.Sprintf("Rectangle with %.2f height and %.2f width", r.h, r.w)
}

func (r rectangle) area() float64 {
	return r.h * r.w
}

func Run() {
	figures := []figure{
		triangle{
			a: 10,
			b: 15,
			c: 13.229,
		},

		rectangle{
			h: 10,
			w: 15,
		},

		circle{
			radius: 18,
		},
	}

	for _, p := range figures {
		printInfo(p)
	}
}

func printInfo(f fmt.Stringer) {
	fmt.Println(f)

	if f, ok := f.(figure); ok {
		fmt.Printf("Area: %.2f\n", f.area())
	}

	if f, ok := f.(triangle); ok {
		fmt.Printf("Triangle's angles are: %v\n", f.angles())
	}
}
