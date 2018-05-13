package flipimage

import (
	"reflect"
	"testing"
)

func TestFlipAndInvertImage(t *testing.T) {
	for _, test := range stringTestCases {
		actual := flipAndInvertImage(test.input)
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("Test [%v], expected [%v], actual [%v]", test.input, test.expected, actual)
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
