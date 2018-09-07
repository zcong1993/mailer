package service

import (
	"github.com/zcong1993/mailer/common"
	"github.com/zcong1993/mailer/mq"
)

// RunService run a mail sender consumer
func RunService(url, qName string, sender common.Sender) {
	conn, ch, msgs := mq.MustDeclareWorker(url, qName)
	defer conn.Close()
	defer ch.Close()

	for msg := range msgs {
		err, requeue := sender.Send(msg.Body)
		if err == nil {
			msg.Ack(false)
			continue
		}
		if err != nil {
			msg.Nack(false, requeue)
		}
	}
}
