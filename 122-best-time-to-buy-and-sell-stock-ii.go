package main

func maxProfit(prices []int) int {
	var profit int
	for i := 0; i < len(prices)-1; i++ {
		j := i + 1
		for ; j < len(prices) && prices[j] >= prices[j-1]; j++ {
		}
		j--
		if j == i {
			continue
		}
		profit += prices[j] - prices[i]
		i = j
	}
	return profit
}
