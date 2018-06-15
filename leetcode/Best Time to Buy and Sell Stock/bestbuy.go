package bestbuy

/**
https://leetcode.com/problems/best-time-to-buy-and-sell-stock/description/
	Say you have an array for which the ith element is the price of a given stock
	on day i.

	If you were only permitted to complete at most one transaction
	(i.e., buy one and sell one share of the stock), design an algorithm to
	find the maximum profit.

	Note that you cannot sell a stock before you buy one.
*/

func maxProfit(prices []int) int {
	if 0 == len(prices) {
		return 0
	}

	var buy, max int

	buy = prices[0]
	for _, p := range prices[1:] {
		if p > buy {
			if (p - buy) > max {
				max = p - buy
			}
		} else {
			buy = p
		}
	}

	return max
}

// Your runtime beats 100.00 % of golang submissions.
