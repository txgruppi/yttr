package yttr_test

import (
	"testing"

	"github.com/txgruppi/yttr"
)

func TestErrorType(t *testing.T) {
	err := yttr.NewError(1, "my test message")
	a := err.Type()
	equal(t, int(1), a)
}

func TestErrorMessage(t *testing.T) {
	err := yttr.NewError(1, "my test message")
	a := err.Message()
	equal(t, "my test message", a)
}

func TestErrorError(t *testing.T) {
	err := yttr.NewError(1, "my test message")
	a := err.Error()
	equal(t, "Remote error: 1 my test message", a)
}
