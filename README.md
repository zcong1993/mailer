# mailer [![Go Report Card](https://goreportcard.com/badge/github.com/zcong1993/mailer)](https://goreportcard.com/report/github.com/zcong1993/mailer)

> mail sender worker in go

## Usage

send msg to queue.

```go
type MailMsg struct {
	From    string   `json:"from"` // required, where the mail send from
	To      []string `json:"to"` // required, where the mail send to
	Subject string   `json:"subject"` // optional, email subject
	Body    string   `json:"body"` // required, email body, plain text or html
	// Tag is not part of email, we use it for analysing
	Tag string `json:"tag"`
}
```

## License

MIT &copy; zcong1993
