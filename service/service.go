package service

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"github.com/zcong1993/amqp-helpers"
	"github.com/zcong1993/debugo"
	"github.com/zcong1993/mailer/common"
	"github.com/zcong1993/mailer/utils"
)

var RouterKeys = []string{"mail"}

var debug = debugo.NewDebug("queue")

// RunService run a mail sender consumer
func RunService(url, exchange, retryExchange string, sender common.Sender, logger common.Logger, maxRetry int) {
	conn := helpers.MustDeclareConn(url)
	ch := helpers.MustDeclareExchange(conn, exchange, nil)
	waitCh := helpers.MustDeclareExchange(conn, retryExchange, nil)

	helpers.MustBindQueue(waitCh, retryExchange, RouterKeys, amqp.Table{"x-dead-letter-exchange": exchange})

	defer conn.Close()
	defer ch.Close()
	defer waitCh.Close()

	_, msgs := helpers.MustDeclareConsumer(ch, exchange, RouterKeys)

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

			logMsg := &common.MailLog{
				MailMsg: m,
				Requeue: requeue,
				Error:   err,
				Retry:   0,
			}
			if !requeue {
				debug.Debugf("not requeue")
				msg.Nack(false, false)
				l <- logMsg
				continue
			}

			p := helpers.CopyMsgToPublishing(msg)

			h := helpers.ParseDeathHeader(p.Headers)

			if h == nil {
				p.Expiration = "3000"
				logMsg.Retry = 1
				debug.Debugf("first retry %s", p.Expiration)
			} else {
				c := h.Count + 1
				logMsg.Retry = c
				if c > maxRetry {
					debug.Debugf("hit max try %d, max is %d", c, maxRetry)
					msg.Nack(false, false)
					l <- logMsg
					continue
				}
				ex := utils.MustToInt(h.OriginalExpiration) * 3
				p.Expiration = fmt.Sprintf("%d", ex)

				debug.Debugf("retry %d - %s", c, p.Expiration)
			}

			waitCh.Publish(retryExchange, msg.RoutingKey, false, false, *p)
			msg.Ack(false)
			l <- logMsg
		}
	}
}
