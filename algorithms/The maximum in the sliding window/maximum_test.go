package main

import (
	"testing"
)

func BenchmarkFlipAndInvertImage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		process(8, []int{2, 7, 3, 1, 5, 2, 6, 2}, 4)
	}
}
