package flipimage

import (
	"reflect"
	"testing"
)

func TestFlipAndInvertImage(t *testing.T) {
	for _, test := range stringTestCases {
		actual := flipAndInvertImage(test.input)
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("\nTest     [%v], \nexpected [%v], \nactual   [%v]", test.input, test.expected, actual)
		}
	}
}

func BenchmarkFlipAndInvertImage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range stringTestCases {
			flipAndInvertImage(test.input)
		}
	}
}
