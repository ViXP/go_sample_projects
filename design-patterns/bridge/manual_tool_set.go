package bridge

import "fmt"

// ManualToolSet is the concrete implementation of ToolSet interface for the manual working processes
type ManualToolSet struct{}

func (s *ManualToolSet) Resaw() {
	fmt.Println("Re-sawing the lumber with a frame saw...")
}

func (s *ManualToolSet) Cut() {
	fmt.Println("Cutting with kataba hand saw...")
}

func (s *ManualToolSet) Drill() {
	fmt.Println("Drilling with the hand drill...")
}

func (s *ManualToolSet) Finish() {
	fmt.Println("Applying finish with a rag and brush...")
}

func (s *ManualToolSet) Joint() {
	fmt.Println("Planing with a jointer plane...")
}

func (s *ManualToolSet) Plane() {
	fmt.Println("Plane to the constant thickness with a scrub & jack planes...")
}

func (s *ManualToolSet) Sand() {
	fmt.Println("Smoothing a surfaces with a smoothing plane and card scraper...")
}

func (s *ManualToolSet) Route() {
	fmt.Println("Process the edges with a chamfer plane...")
}

// Explicit check for interface conformity
var _ ToolSet = &ManualToolSet{}
