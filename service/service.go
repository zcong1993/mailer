package service

import (
	"encoding/json"
	"github.com/zcong1993/mailer/common"
	"github.com/zcong1993/mailer/mq"
)

// RunService run a mail sender consumer
func RunService(url, qName string, sender common.Sender, logger common.Logger) {
	conn, ch, msgs := mq.MustDeclareWorker(url, qName)
	defer conn.Close()
	defer ch.Close()

	l := logger.GetChannel()

	for msg := range msgs {
		var m common.MailMsg
		err := json.Unmarshal(msg.Body, &m)
		if err != nil {
			l <- &common.MailLog{
				MailMsg: m,
				Error:   err,
			}
			continue
		}
		err, requeue := sender.Send(m)
		if err == nil {
			msg.Ack(false)
			l <- &common.MailLog{
				MailMsg: m,
			}
			continue
		}
		if err != nil {
			l <- &common.MailLog{
				MailMsg: m,
				Requeue: requeue,
				Error:   err,
			}
			msg.Nack(false, requeue)
		}
	}
}
