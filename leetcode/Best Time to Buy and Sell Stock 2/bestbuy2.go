package bestbuy

/**
https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/description/
	Say you have an array for which the ith element is the price of a given stock
	on day i.

	Design an algorithm to find the maximum profit. You may complete as many
	transactions as you like (i.e., buy one and sell one share of the stock
	multiple times).

	Note: You may not engage in multiple transactions at the same time
	(i.e., you must sell the stock before you buy again).
*/

func maxProfit(prices []int) int {
	if 0 == len(prices) {
		return 0
	}

	var buy, max, current int

	buy = prices[0]
	for i, p := range prices[1:] {
		if p > buy {
			current = p - buy

			if i < len(prices)-1 {
				buy = p
			}
		} else {
			buy = p
		}
		max += current
		current = 0
	}

	return max
}

// Your runtime beats 97.85 % of golang submissions.
