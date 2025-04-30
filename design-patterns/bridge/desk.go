package bridge

import "fmt"

// Desk is the specific Project that can be built through the bridge to the tool sets
type Desk struct {
	Project
}

// Build implements the building algorithm
func (d *Desk) Build() {
	d.CheckTools()
	d.Tools.Joint()
	d.Tools.Plane()
	d.Tools.Cut()
	d.Tools.Route()
	d.Tools.Sand()
	fmt.Printf("The desk is built!\n\n")
}

// Explicit check for interface conformity
var _ ProjectBuild = &Desk{}
