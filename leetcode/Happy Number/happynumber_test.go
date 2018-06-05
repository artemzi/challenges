package happynumber

import (
	"testing"
)

func TestIsHappy(t *testing.T) {
	expect := true
	actual := isHappy(19)
	if actual != expect {
		t.Errorf("Wrong result for case #1. expect: %v, got: %v\n", expect, actual)
	}
}
