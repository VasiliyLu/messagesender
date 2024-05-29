package smsru_test

import (
	"os"
	"testing"

	"tgbotokna/pkg/messagesender"
	"tgbotokna/pkg/messagesender/smsru"
	"tgbotokna/pkg/testingex"
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
