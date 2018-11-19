package service

import (
	"encoding/json"
	"github.com/zcong1993/amqp-retry-worker"
	"github.com/zcong1993/debugo"
	"github.com/zcong1993/mailer/common"
	"gopkg.in/go-playground/validator.v9"
)

var RouterKeys = []string{""}

var debug = debugo.NewDebug("queue")
var validate = validator.New()

type Worker struct {
	Sender common.Sender
	Logger common.Logger
}

func (w *Worker) Do(payload []byte, routerKey string) (error, bool) {
	l := w.Logger.GetChannel()
	var msg common.MailMsg
	err := json.Unmarshal(payload, &msg)
	if err != nil {
		l <- &common.MailLog{
			MailMsg: msg,
			Error:   err,
		}
		return err, false
	}
	err = validate.Struct(msg)
	if err != nil {
		l <- &common.MailLog{
			MailMsg: msg,
			Error:   err,
		}
		return err, false
	}
	err, retry := w.Sender.Send(msg)
	if err != nil {
		logMsg := &common.MailLog{
			MailMsg: msg,
			Error:   err,
		}
		l <- logMsg
	} else {
		l <- &common.MailLog{
			MailMsg: msg,
		}
	}

	return err, retry
}

// RunService run a mail sender consumer
func RunService(url, exchange, retryExchange, qName string, sender common.Sender, logger common.Logger, maxRetry int) {
	bf := worker.NewDefaultBackoff(10000)
	w := worker.NewRetryWorker(url, exchange, qName, RouterKeys, &Worker{Sender: sender, Logger: logger}, bf, maxRetry, false)
	w.Run()
}
