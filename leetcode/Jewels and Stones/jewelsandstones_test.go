package jewelsandstones

import "testing"

func TestNumJewelsInStones(t *testing.T) {
	expect := 3
	actual := numJewelsInStones("aA", "aAAbbbb")
	if actual != expect {
		t.Errorf("Wrong result for case #1. expect: %v, got: %v\n", expect, actual)
	}

	expect = 0
	actual = numJewelsInStones("z", "ZZ")
	if actual != expect {
		t.Errorf("Wrong result for case #2. expect: %v, got: %v\n", expect, actual)
	}
}
