package messagesender

type Message struct {
	Contact string
	Subject string
	Text    string
}

type MessageSender interface {
	Send(message *Message) error
}
