package client

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/zcong1993/mailer/common"
	"math/rand"
)

// MockSender is mock sender implement
type MockSender struct{}

// Send impl the Sender interface
func (ms *MockSender) Send(mail common.MailMsg) (error, bool) {
	nn := rand.Intn(10)
	if nn < 2 {
		fmt.Printf("success handle work %v\n", mail)
		return nil, false
	}

	return errors.New("mock failed"), true
}
