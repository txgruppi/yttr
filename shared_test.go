package yttr_test

import "testing"

func equal(t *testing.T, expected, actual interface{}) {
	if expected != actual {
		t.Errorf(
			"Expected %T(%v) but got %T(%v)",
			expected,
			expected,
			actual,
			actual,
		)
	}
}

func notEqual(t *testing.T, expected, actual interface{}) {
	if expected == actual {
		t.Errorf(
			"Expected not %T(%v)",
			expected,
			expected,
			actual,
			actual,
		)
	}
}
