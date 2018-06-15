package bestbuy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	input  []int
	expect int
}{

	{
		[]int{7, 1, 5, 3, 6, 4},
		5,
	},
	{
		[]int{7, 6, 4, 3, 1},
		0,
	},
}

func Test_coinChange(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.expect, maxProfit(tc.input), "input:%v", tc)
	}
}

func Benchmark_coinChange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			maxProfit(tc.input)
		}
	}
}
