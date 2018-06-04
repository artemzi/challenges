package intersection

import (
	"reflect"
	"testing"
)

func TestIntersection(t *testing.T) {
	expect := []int{2}
	actual := intersection([]int{1, 2, 2, 1}, []int{2, 2})
	if !reflect.DeepEqual(actual, expect) {
		t.Errorf("Wrong result for case #1. expect: %v, got: %v\n", expect, actual)
	}
}
