package event

type Payload struct {
	Name       string `json:"name"`
	Data       string `json:"data"`
	RoutingKey string
}

func (p *Payload) Handle() {
	var handler Handler

	switch p.RoutingKey {
	case "log.INFO":
		handler = &LogHandler{}
	case "auth.INFO":
		handler = &AuthHandler{}
	default:
		handler = &DefaultHandler{}
	}
	handler.Handle(p)
}
