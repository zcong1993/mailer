package common

// DEFAULT_TAG is name of default tag
const DEFAULT_TAG = "default"

// Sender is interface of mail sender
type Sender interface {
	Send(mail MailMsg) (error, bool)
}

// MailMsg is msg we received from rabbit mq
type MailMsg struct {
	From    string   `json:"from" validate:"required,email"`
	To      []string `json:"to" validate:"required,dive,email"`
	Subject string   `json:"subject"`
	Body    string   `json:"body" validate:"required"`
	// Tag is not part of email, we use it for analysing
	Tag string `json:"tag"`
	// ID is uuid of this message
	ID string `json:"id" validate:"required"`
}

// MailLog is struct of mail log
type MailLog struct {
	MailMsg
	Requeue bool
	Error   error
	Retry   int
}

// Logger is logger interface for service log
type Logger interface {
	GetChannel() chan<- *MailLog
}
