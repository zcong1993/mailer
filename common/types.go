package common

// Sender is interface of mail sender
type Sender interface {
	Send([]byte) (error, bool)
}
