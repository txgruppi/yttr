package yttr_test

import (
	"testing"

	"github.com/txgruppi/yttr"
)

func TestBoolString(t *testing.T) {
	b0 := yttr.Bool(true)
	b1 := yttr.Bool(false)

	equal(t, "true", b0.String())
	equal(t, "false", b1.String())
}
