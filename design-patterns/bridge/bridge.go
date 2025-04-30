// Package bridge implements the Bridge design pattern.
// In the context of woodworking projects, the bridge is implemented through dependency inversion.
// The specific projects (Desk, Chair) are not tightly coupled to a particular tool set (ManualToolSet or PowerToolSet),
// but instead depend on an abstract interface (ToolSet).
// Each project conforms to the ProjectBuild interface and inherits from the Project struct,
// which encapsulates the abstract ToolSet and validation method, allowing for flexible tool swapping without modifying
// the project logic.
package bridge

func Run() {
	// Defining the manual and power tools
	manualTools := ManualToolSet{}
	powerTools := PowerToolSet{}

	// Building the chair with a manual tools
	chairProject := Chair{Project{&manualTools}}
	chairProject.Build()

	// Building it with a power tools
	chairProject.Tools = &powerTools
	chairProject.Build()

	// Building the desk with a power tools
	deskProject := Desk{Project{&powerTools}}
	deskProject.Build()

	// Try to build the project without toolset
	wrongProject := Chair{}
	wrongProject.Build()
}
