package common

// Sender is interface of mail sender
type Sender interface {
	Send([]byte) (error, bool)
}

// MailMsg is msg we received from rabbit mq
type MailMsg struct {
	From    string   `json:"from"`
	To      []string `json:"to"`
	Subject string   `json:"subject"`
	Body    string   `json:"body"`
	// Tag is not part of email, we use it for analysing
	Tag string `json:"tag"`
	// ID is uuid of this message
	ID string `json:"id"`
}
