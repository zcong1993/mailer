package main

import (
	"github.com/zcong1993/mailer/client"
	"github.com/zcong1993/mailer/service"
	"github.com/zcong1993/mailer/utils"
)

func main() {
	rabbit := utils.EnvOrDefault("RABBIT", "amqp://guest:guest@localhost:5672/")
	qName := "mail"
	mockSender := &client.MockSender{}
	service.RunService(rabbit, qName, mockSender)
}
