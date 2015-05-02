package yttr

import "fmt"

func NewError(code int, message string) Error {
	return &rerror{
		code:    code,
		message: message,
	}
}

type Error interface {
	Error() string
	Type() int
	Message() string
}

type rerror struct {
	code    int
	message string
}

func (e *rerror) Type() int {
	return e.code
}

func (e *rerror) Message() string {
	return e.message
}

func (e *rerror) Error() string {
	return fmt.Sprintf("Remote error: %d %s", e.code, e.message)
}
