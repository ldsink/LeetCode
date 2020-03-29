package main

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	var result int
	for i, j := len(prices)-2, prices[len(prices)-1]; i >= 0; i-- {
		if j-prices[i] > result {
			result = j - prices[i]
		}
		if j < prices[i] {
			j = prices[i]
		}
	}
	return result
}
