package main

func maxProfit(k int, prices []int) int {
	if k == 0 || len(prices) < 2 {
		return 0
	}

	if k > len(prices)/2 {
		k = len(prices) / 2
		var result int
		for i := 1; i < len(prices); i++ {
			if prices[i] > prices[i-1] {
				result += prices[i] - prices[i-1]
			}
		}
		return result
	}

	costPrice := [2][]int{make([]int, 2*k), make([]int, 2*k)}
	profit := [2][]int{make([]int, 2*k), make([]int, 2*k)}
	for i := 0; i < 2*k; i++ {
		costPrice[0][i] = -1
		costPrice[1][i] = -1
	}
	for i := 0; i < len(prices); i++ {
		today := i % 2
		yesterday := today ^ 1
		// buy condition
		costPrice[today][0] = costPrice[yesterday][0]
		if costPrice[today][0] < 0 || costPrice[today][0] > prices[i] {
			costPrice[today][0] = prices[i]
		}
		for j := 2; j < 2*k; j += 2 {
			costPrice[today][j] = costPrice[yesterday][j]
			profit[today][j] = profit[yesterday][j]
			if costPrice[today][j] < 0 || costPrice[today][j] > prices[i] {
				costPrice[today][j] = prices[i]
			}
			if profit[today][j] < profit[yesterday][j-1]-(prices[i]-costPrice[today][j]) {
				profit[today][j] = profit[yesterday][j-1]
				costPrice[today][j] = prices[i]
			}
		}
		// sell condition
		for j := 1; j < 2*k; j += 2 {
			costPrice[today][j] = costPrice[yesterday][j]
			profit[today][j] = profit[yesterday][j]
			if costPrice[yesterday][j-1] >= 0 && profit[today][j] < profit[yesterday][j-1]+prices[i]-costPrice[yesterday][j-1] {
				profit[today][j] = profit[yesterday][j-1] + prices[i] - costPrice[yesterday][j-1]
			}
		}
	}
	return profit[(len(prices)-1)%2][2*k-1]
}
