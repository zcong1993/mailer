package client

import (
	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	"github.com/zcong1993/mailer/common"
	"os"
)

// DefaultLogger is default logger struct
type DefaultLogger struct {
	ch chan<- *common.MailLog
}

// NewDefaultLogger construct a new default logger instance
func NewDefaultLogger(buffer int) *DefaultLogger {
	log.SetHandler(cli.New(os.Stdout))
	ch := make(chan *common.MailLog, buffer)
	go func() {
		for msg := range ch {
			if msg.Error == nil {
				l := log.WithFields(log.Fields{
					"ID":  msg.ID,
					"To":  msg.To,
					"Tag": msg.Tag,
				})
				l.Infof("success")
			} else {
				l := log.WithFields(log.Fields{
					"ID":      msg.ID,
					"To":      msg.To,
					"Tag":     msg.Tag,
					"Requeue": msg.Requeue,
					"Retry":   msg.Retry,
				})
				l.Error(msg.Error.Error())
			}
		}
	}()
	return &DefaultLogger{
		ch: ch,
	}
}

// GetChannel impl logger interface
func (l *DefaultLogger) GetChannel() chan<- *common.MailLog {
	return l.ch
}
