package goatlatin

import (
	"testing"
)

func TestToGoatLatin(t *testing.T) {
	for _, test := range stringTestCases {
		actual := toGoatLatin(test.input)
		if actual != test.expected {
			t.Errorf("Test [%s], expected [%s], actual [%s]", test.input, test.expected, actual)
		}
	}
}

func BenchmarkToGoatLatin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range stringTestCases {
			toGoatLatin(test.input)
		}
	}
}
