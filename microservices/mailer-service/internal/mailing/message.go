package mailing

type Message struct {
	From        string
	FromName    string
	To          string
	Subject     string
	Body        string
	Attachments []string
	DataMap     map[string]any
}

func NewMessage(to, subject, body string) *Message {
	msg := Message{
		To:      to,
		Subject: subject,
		Body:    body,
		DataMap: make(map[string]any),
	}
	msg.DataMap = map[string]any{"message": msg.Body}

	return &msg
}
