package smsru_test

import (
	"os"
	"testing"

	"github.com/VasiliyLu/messagesender"
	"github.com/VasiliyLu/messagesender/smsru"
	"github.com/VasiliyLu/messagesender/pkg/testingex"
)

func TestSendSms(t *testing.T) {
	if testingex.CheckIntegrationEnv(t) {
		return
	}
	//arrange
	contact := os.Getenv("SMSRU_TEST_PHONE")
	smsRuConfig := smsru.SmsRuConfig{
		Token: os.Getenv("SMSRU_TOKEN"),
		Test:  false,
	}
	smsRu := smsru.NewSmsRu(&smsRuConfig)
	//act
	err := smsRu.Send(&messagesender.Message{
		Contact: contact,
		Text:    "Test sms"})
	//assert
	if err != nil {
		t.Errorf(err.Error())
	}
}
