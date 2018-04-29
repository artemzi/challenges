package friendsages

import (
	"testing"
)

func TestNumFriendRequests(t *testing.T) {
	for _, test := range stringTestCases {
		actual := numFriendRequests(test.input)
		if actual != test.expected {
			t.Errorf("Test [%v], expected [%v], actual [%v]", test.input, test.expected, actual)
		}
	}
}

func BenchmarkNumFriendRequests(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range stringTestCases {
			numFriendRequests(test.input)
		}
	}
}
