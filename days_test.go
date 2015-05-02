package yttr_test

import (
	"testing"

	"github.com/txgruppi/yttr"
)

func TestDaysString(t *testing.T) {
	d0 := yttr.Days(-1)
	d1 := yttr.Days(0)
	d2 := yttr.Days(1)

	equal(t, "-1", d0.String())
	equal(t, "0", d1.String())
	equal(t, "1", d2.String())
}
