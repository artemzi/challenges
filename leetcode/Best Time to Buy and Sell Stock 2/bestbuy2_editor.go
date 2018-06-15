package bestbuy

/**
Better solution from other leetcode user
*/

func maxProfit2(prices []int) int {
	profit := 0
	for i := 0; i+1 < len(prices); i++ {
		if prices[i] < prices[i+1] {
			profit += (prices[i+1] - prices[i])
		}
	}
	return profit
}

// goos: linux
// goarch: amd64
// Benchmark_maxProfit    	50000000	        32.5 ns/op
// Benchmark_maxProfit2   	20000000	        57.2 ns/op
