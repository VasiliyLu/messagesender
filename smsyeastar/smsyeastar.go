package smsyeastar

import (
	"errors"
	"fmt"
	"github.com/nyaruka/phonenumbers"
	"io"
	"github.com/VasiliyLu/messagesender"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type SmsYeastarConfig struct {
	Host    string `yaml:"host"`
	ApiUser string `yaml:"apiuser"`
	ApiPass string `yaml:"apipass"`
	Port    int    `yaml:"port"`
}

type SmsYeastar struct {
	smsYeastarConfig *SmsYeastarConfig
}

func formatPhone(phone string) (string, error) {
	phonenumber, err := phonenumbers.Parse(phone, "RU")
	if err != nil {
		return "", err
	}
	return "+" + strconv.Itoa(int(*phonenumber.CountryCode)) + strconv.FormatUint(*phonenumber.NationalNumber, 10), nil
}

// Send implements messagesender.MessageSender.
func (s SmsYeastar) Send(message *messagesender.Message) error {

	contact, err := formatPhone(message.Contact)
	if err != nil {
		return err
	}
	q := url.Values{}
	q.Add("account", s.smsYeastarConfig.ApiUser)
	q.Add("password", s.smsYeastarConfig.ApiPass)
	q.Add("port", strconv.Itoa(s.smsYeastarConfig.Port))

	q.Add("destination", contact)
	q.Add("content", message.Text)

	url := fmt.Sprintf("http://%s/cgi/WebCGI?1500101=%s", s.smsYeastarConfig.Host, q.Encode())

	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if !strings.Contains(string(bodyBytes), "success") {
		return errors.New("Error on sending message: " + string(bodyBytes))
	}

	return nil
}

func New(smsYeastartConfig *SmsYeastarConfig) messagesender.MessageSender {
	return &SmsYeastar{smsYeastarConfig: smsYeastartConfig}
}
