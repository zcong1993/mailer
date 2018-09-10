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
			l := log.WithFields(log.Fields{
				"ID":  msg.ID,
				"To":  msg.To,
				"Tag": msg.Tag,
			})
			if msg.Error == nil {
				l.Infof("success")
			} else {
				l.WithField("requeue", msg.Requeue)
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
