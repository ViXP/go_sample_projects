package bridge

// ToolSet defines a set of woodworking operations regardless of tool type (manual, power etc.).
type ToolSet interface {
	Cut()
	Resaw()
	Joint()
	Plane()
	Drill()
	Route()
	Sand()
	Finish()
}
