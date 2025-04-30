package bridge

// ProjectBuild defines the abstraction for any woodworking project that can be built using a ToolSet.
type ProjectBuild interface {
	Build()
}

// Project is the abstract Project type that stores the specific ToolSet field
type Project struct {
	Tools ToolSet
}

// CheckTools it the validation method for the Tools presence
func (p *Project) CheckTools() {
	if p.Tools == nil {
		panic("toolset is missing")
	}
}
