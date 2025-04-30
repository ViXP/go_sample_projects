package bridge

import "fmt"

// PowerToolSet is the concrete implementation of ToolSet interface for the electric instruments operated processes
type PowerToolSet struct{}

func (s *PowerToolSet) Resaw() {
	fmt.Println("Re-sawing the lumber with a band saw...")
}

func (s *PowerToolSet) Cut() {
	fmt.Println("Cutting with a table saw and miter saw...")
}

func (s *PowerToolSet) Drill() {
	fmt.Println("Drilling with the drill press...")
}

func (s *PowerToolSet) Finish() {
	fmt.Println("Applying finish with a spray gun...")
}

func (s *PowerToolSet) Joint() {
	fmt.Println("Planing with a jointer...")
}

func (s *PowerToolSet) Plane() {
	fmt.Println("Plane to the constant thickness with a thickness planer...")
}

func (s *PowerToolSet) Sand() {
	fmt.Println("Sanding the surface with the orbital sander...")
}

func (s *PowerToolSet) Route() {
	fmt.Println("Process the edges with a router...")
}

// Explicit check for interface conformity
var _ ToolSet = &PowerToolSet{}
