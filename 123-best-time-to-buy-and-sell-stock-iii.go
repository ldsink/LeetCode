package main

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	leftMax := make([]int, len(prices))
	for i, j := 1, prices[0]; i < len(prices); i++ {
		leftMax[i] = leftMax[i-1]
		if leftMax[i] < prices[i]-j {
			leftMax[i] = prices[i] - j
		}
		if j > prices[i] {
			j = prices[i]
		}
	}
	rightMax := make([]int, len(prices))
	for i, j := len(prices)-2, prices[len(prices)-1]; i >= 0; i-- {
		rightMax[i] = rightMax[i+1]
		if rightMax[i] < j-prices[i] {
			rightMax[i] = j - prices[i]
		}
		if j < prices[i] {
			j = prices[i]
		}
	}
	var result int
	for i := 0; i < len(prices); i++ {
		r := leftMax[i]
		if i+1 < len(prices) {
			r += rightMax[i+1]
		}
		if result < r {
			result = r
		}
	}
	return result
}
