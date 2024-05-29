package smtpsender_test

import (
	"testing"
	"tgbotokna/pkg/messagesender"
	"tgbotokna/pkg/messagesender/smtpsender"
	"tgbotokna/pkg/testingex"
)

func TestSendTLS(t *testing.T) {
	if testingex.CheckIntegrationEnv(t) {
		return
	}
	//arrange
	config := smtpsender.SmtpConfig{
		From:     "from@example.com",
		Host:     "sandbox.smtp.mailtrap.io",
		Port:     "465",
		Username: "64a52dfd714f86",
		Password: "bc24f71379c0ff",
	}
	smtpSender := smtpsender.New(&config)
	msg := messagesender.Message{
		Contact: "to@example.com",
		Subject: "Test subject",
		Text:    "Test Message",
	}
	//act
	err := smtpSender.Send(&msg)
	//assert
	if err != nil {
		t.Errorf(err.Error())
	}
}
