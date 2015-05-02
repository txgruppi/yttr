package yttr_test

import (
	"testing"

	"github.com/txgruppi/yttr"
)

func TestSizeString(t *testing.T) {
	s0 := yttr.Size(-1)
	s1 := yttr.Size(0)
	s2 := yttr.Size(1)

	equal(t, "-1", s0.String())
	equal(t, "0", s1.String())
	equal(t, "1", s2.String())
}
