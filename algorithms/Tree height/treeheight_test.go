package main

import (
	"testing"
)

func TestFlipAndInvertImage(t *testing.T) {
	for _, test := range stringTestCases {
		actual := getMaxDeph(test.input.val)
		if actual != test.expected {
			t.Errorf("\nTest     [%v], \nexpected [%v], \nactual   [%v]", test.input, test.expected, actual)
		}
	}
}

func BenchmarkFlipAndInvertImage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range stringTestCases {
			getMaxDeph(test.input.val)
		}
	}
}
