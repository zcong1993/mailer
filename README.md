# mailer [![Go Report Card](https://goreportcard.com/badge/github.com/zcong1993/mailer)](https://goreportcard.com/report/github.com/zcong1993/mailer)

> mail sender worker in go

## Usage

send msg to queue.

```go
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
```

## License

MIT &copy; zcong1993
