package client

import (
	"encoding/json"
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
func (ms *MailSender) Send(msg []byte) (error, bool) {
	m, data, err := newMail(msg)
	if err != nil {
		logger.Error(err, data)
		return errors.New("invalid mail. "), false
	}
	err = mail.Send(ms.s, m)
	if err != nil {
		// TODO: make sure which should retry
		logger.Error(err, data)
		return err, false
	}
	logger.Infof("send mail %+v\n", data)
	return nil, false
}

func newMail(msg []byte) (*mail.Message, common.MailMsg, error) {
	var data common.MailMsg
	err := json.Unmarshal(msg, &data)
	if err != nil {
		return nil, data, err
	}

	if data.From == "" {
		return nil, data, errors.New("from to is required. ")
	}

	if len(data.To) == 0 {
		return nil, data, errors.New("send to is required. ")
	}

	if data.Body == "" {
		return nil, data, errors.New("send empty to user? ")
	}

	m := mail.NewMessage()

	m.SetBody("text/html", data.Body)
	m.SetHeader("From", data.From)
	m.SetHeader("To", data.To...)
	m.SetHeader("Subject", data.Subject)

	return m, data, nil
}
