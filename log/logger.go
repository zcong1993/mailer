package log

import (
	"github.com/sirupsen/logrus"
	"github.com/zcong1993/mailer/utils"
)

// Logger is logger instance
var Logger *logrus.Logger

func init() {
	l := logrus.New()
	env := utils.EnvOrDefault("LOG_ENV", "debug")
	if env == "debug" {
		l.SetLevel(logrus.DebugLevel)
	}
	Logger = l
}
