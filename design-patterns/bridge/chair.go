package bridge

import "fmt"

// Chair is the specific Project that can be built through the bridge to the tool sets
type Chair struct {
	Project
}

// Build implements the building algorithm
func (c *Chair) Build() {
	c.CheckTools()
	c.Tools.Resaw()
	c.Tools.Cut()
	c.Tools.Joint()
	c.Tools.Plane()
	c.Tools.Route()
	c.Tools.Sand()
	c.Tools.Finish()
	fmt.Printf("The chair is built!\n\n")
}

// Explicit check for interface conformity
var _ ProjectBuild = &Chair{}
