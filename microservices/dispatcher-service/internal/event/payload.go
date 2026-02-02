package event

type Payload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func (p *Payload) Handle() {
	var handler Handler

	switch p.Name {
	case "log":
		handler = &LogHandler{}
	case "auth":
		handler = &AuthHandler{}
	default:
		handler = &DefaultHandler{}
	}
	handler.Handle(p)
}
