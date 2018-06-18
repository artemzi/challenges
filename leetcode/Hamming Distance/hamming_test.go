package hamming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	x      int
	y      int
	expect int
}{

	{
		1,
		4,
		2,
	},
	{
		680142203,
		1111953568,
		19,
	},
}

func Test_maxProfit(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.expect, hammingDistance(tc.x, tc.y), "input:%v", tc)
	}
}

func Benchmark_maxProfit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			hammingDistance(tc.x, tc.y)
		}
	}
}
