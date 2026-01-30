package server

type ProxyRequestPayload struct {
	Action string        `json:"action"`
	Auth   AuthPayload   `json:"auth,omitempty"`
	Log    LogPayload    `json:"log,omitempty"`
	Mail   MailerPayload `json:"mail,omitempty"`
}

type AuthPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LogPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

type MailerPayload struct {
	To       string `json:"to"`
	From     string `json:"from,omitempty"`
	FromName string `json:"from_name,omitempty"`
	Subject  string `json:"subject"`
	Body     string `json:"body"`
}
