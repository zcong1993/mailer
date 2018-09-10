package main

import (
	"github.com/zcong1993/mailer/client"
	"github.com/zcong1993/mailer/service"
	"github.com/zcong1993/mailer/utils"
)

func main() {
	rabbit := utils.EnvOrDefault("RABBIT", "amqp://guest:guest@localhost:5672/")
	qName := utils.EnvOrDefault("QUEUE_NAME", "mail")
	sender := client.MustNewMailSender("smtp.gmail.com", 465, utils.RequiredEnv("MAIL_ACCOUNT"), utils.RequiredEnv("MAIL_PASSWORD"))
	logger := client.NewDefaultLogger(10)
	//mongoLogger := client.NewMongoLogger("", "mailer", "log", 10)
	service.RunService(rabbit, qName, sender, logger)
}
