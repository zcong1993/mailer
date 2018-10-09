package main

import (
	"github.com/zcong1993/mailer/client"
	"github.com/zcong1993/mailer/service"
	"github.com/zcong1993/mailer/utils"
)

func main() {
	rabbit := utils.EnvOrDefault("RABBIT", "amqp://guest:guest@localhost:5672/")
	exchangeName := utils.EnvOrDefault("EXCHANGE_NAME", "mail")
	retryExchangeName := utils.EnvOrDefault("RETRY_EXCHANGE_NAME", "mail_retry")
	qName := utils.EnvOrDefault("QUEUE_NAME", "mail")
	sender := client.MustNewMailSender(utils.EnvOrDefault("SMTP_HOST", "smtp.gmail.com"), utils.MustToInt(utils.EnvOrDefault("SMTP_PORT", "465")), utils.RequiredEnv("MAIL_ACCOUNT"), utils.RequiredEnv("MAIL_PASSWORD"))
	logger := client.NewDefaultLogger(10)
	//mockSender := &client.MockSender{}
	service.RunService(rabbit, exchangeName, retryExchangeName, qName, sender, logger, 3)
}
