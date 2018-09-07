package client

import (
	"fmt"
	"github.com/pkg/errors"
	"math/rand"
)

// MockSender is mock sender implement
type MockSender struct{}

// Send impl the Sender interface
func (ms *MockSender) Send(msg []byte) (error, bool) {
	n := string(msg)
	println(n)
	nn := rand.Intn(10)
	if nn < 5 {
		fmt.Printf("success handle work %s\n", n)
		return nil, false
	}

	return errors.New("mock failed"), true
}
