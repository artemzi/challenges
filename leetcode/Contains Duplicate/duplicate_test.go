package duplicate

import "testing"

func TestContainsDuplicate(t *testing.T) {
	actual := containsDuplicate([]int{1, 2, 3, 1})
	if actual != true {
		t.Errorf("Wrong result for case #1. expect: %v, got: %v\n", true, actual)
	}
}
