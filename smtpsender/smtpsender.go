package smtpsender

import (
	"bytes"
	"fmt"
	"github.com/VasiliyLu/messagesender"
	"net/smtp"
	"strings"
)

type SmtpConfig struct {
	From     string `yaml:"from"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
}

type SmtpSender struct {
	smtpConfig *SmtpConfig
}

// Send implements messagesender.MessageSender.
func (s *SmtpSender) Send(message *messagesender.Message) error {

	// Receiver email address.
	to := []string{
		strings.ToLower(message.Contact),
	}

	// Authentication.
	auth := smtp.PlainAuth("", s.smtpConfig.Username, s.smtpConfig.Password, s.smtpConfig.Host)

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: %s \n%s\n\n", strings.Replace(message.Subject, "\n", " ", -1), mimeHeaders)))
	body.WriteString(message.Text)

	// Sending email.
	err := smtp.SendMail(s.smtpConfig.Host+":"+s.smtpConfig.Port, auth, s.smtpConfig.From, to, body.Bytes())
	if err != nil {
		return err
	}
	return nil
}

func New(smtpConfig *SmtpConfig) messagesender.MessageSender {
	return &SmtpSender{smtpConfig: smtpConfig}
}
