package main

func _maxProfit(prices []int) int {
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

func maxProfit(prices []int) int {
	var result int
	for i := 0; i < len(prices); i++ {
		if r := _maxProfit(prices[:i]) + _maxProfit(prices[i:]); result < r {
			result = r
		}
	}
	return result
}
