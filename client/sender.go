package client

import (
	"errors"
	"github.com/go-mail/mail"
	"github.com/zcong1993/mailer/common"
	"github.com/zcong1993/mailer/log"
)

var logger = log.Logger

// MailSender is true mail sender
type MailSender struct {
	d *mail.Dialer
	s mail.SendCloser
}

// MustNewMailSender return a new MailSender
func MustNewMailSender(host string, port int, username, password string) *MailSender {
	d := mail.NewDialer(host, port, username, password)
	s, err := d.Dial()
	if err != nil {
		panic(err)
	}
	return &MailSender{
		d: d,
		s: s,
	}
}

// Send impl mail sender
func (ms *MailSender) Send(msg common.MailMsg) (error, bool) {
	m, err := newMail(msg)
	if err != nil {
		logger.Error(err, msg)
		return err, false
	}
	err = mail.Send(ms.s, m)
	if err != nil {
		// TODO: make sure which should retry
		return err, false
	}
	return nil, false
}

func newMail(data common.MailMsg) (*mail.Message, error) {
	if data.From == "" {
		return nil, errors.New("from to is required. ")
	}

	if len(data.To) == 0 {
		return nil, errors.New("send to is required. ")
	}

	if data.Body == "" {
		return nil, errors.New("send empty to user? ")
	}

	m := mail.NewMessage()

	m.SetBody("text/html", data.Body)
	m.SetHeader("From", data.From)
	m.SetHeader("To", data.To...)
	m.SetHeader("Subject", data.Subject)

	return m, nil
}
