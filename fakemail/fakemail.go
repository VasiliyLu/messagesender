package fakemail

import (
	"fmt"
	"github.com/VasiliyLu/messagesender"
)

type FakeMail struct {
}

// Send implements messagesender.MessageSender.
func (*FakeMail) Send(message *messagesender.Message) error {
	fmt.Printf(message.Text)
	return nil
}

func New() messagesender.MessageSender {
	return &FakeMail{}
}
