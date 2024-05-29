package smsru

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"tgbotokna/pkg/messagesender"

	"github.com/samber/lo"
	"golang.org/x/exp/maps"
)

type SmsRuConfig struct {
	Token string `yaml:"token"`
	Test  bool   `yaml:"test"`
}

type SmsRu struct {
	smsRuConfig *SmsRuConfig
}

type ResponseSMSStatus struct {
	Status     string `json:"status"`
	StatusCode int    `json:"status_code"`
	SMSID      string `json:"sms_id,omitempty"`
	StatusText string `json:"status_text"`
}

type Response struct {
	Status     string                       `json:"status"`
	StatusCode int                          `json:"status_code"`
	SMS        map[string]ResponseSMSStatus `json:"sms"`
	Balance    float64                      `json:"balance"`
}

// Send implements messagesender.MessageSender.
func (s SmsRu) Send(message *messagesender.Message) error {

	q := url.Values{}
	q.Add("api_id", s.smsRuConfig.Token)
	q.Add("to", message.Contact)
	q.Add("msg", message.Text)
	q.Add("json", "1")
	if s.smsRuConfig.Test {
		q.Add("test", "1")
	}

	resp, err := http.Post("https://sms.ru/sms/send?"+q.Encode(), "", nil)
	if err != nil {
		return err
	}

	var responseDto Response
	data, _ := io.ReadAll(resp.Body)
	err = json.Unmarshal(data, &responseDto)
	if err != nil {
		return err
	}

	if responseDto.StatusCode != 100 {
		return errors.New("SMS.RU: " + responseDto.Status)
	}

	smsStatus, found := lo.Find[ResponseSMSStatus](maps.Values(responseDto.SMS), func(item ResponseSMSStatus) bool {
		return item.StatusCode != 100
	})
	if found {
		return errors.New("SMS.RU: " + smsStatus.StatusText)
	}

	return nil
}

func NewSmsRu(smsRuConfig *SmsRuConfig) messagesender.MessageSender {
	return &SmsRu{smsRuConfig: smsRuConfig}
}
