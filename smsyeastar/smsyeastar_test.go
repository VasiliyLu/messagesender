package smsyeastar_test

import (
	"os"
	"strconv"
	"testing"
	"tgbotokna/pkg/messagesender"
	"tgbotokna/pkg/messagesender/smsyeastar"
	"tgbotokna/pkg/testingex"
)

var testCases = []string{"Ваш код: 7705", "Добро пожаловать!", "Тестовое сообщение"}

func TestSendSms(t *testing.T) {
	if testingex.CheckIntegrationEnv(t) {
		return
	}
	//arrange
	contact := os.Getenv("SMSRU_TEST_PHONE")
	port, _ := strconv.Atoi(os.Getenv("YEA_PORT"))
	smsYeaConfig := smsyeastar.SmsYeastarConfig{
		Host:    os.Getenv("YEA_HOST"),
		Port:    port,
		ApiUser: os.Getenv("YEA_USER"),
		ApiPass: os.Getenv("YEA_PASS"),
	}
	smsYea := smsyeastar.New(&smsYeaConfig)
	for _, message := range testCases {
		//act
		err := smsYea.Send(&messagesender.Message{
			Contact: contact,
			Text:    message})
		//assert
		if err != nil {
			t.Errorf(err.Error())
		}
	}

}
