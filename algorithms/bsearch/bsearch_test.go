package bsearch

import (
	"testing"
)

func TestBsearch(t *testing.T) {
	for _, test := range stringTestCases {
		actual := Bsearch(test.input.data, test.input.val)
		if actual != test.expected {
			t.Errorf("\nTest     [%v], \nexpected [%v], \nactual   [%v]", test.input, test.expected, actual)
		}
	}
}

func BenchmarkBsearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range stringTestCases {
			Bsearch(test.input.data, test.input.val)
		}
	}
}
