package main

func dp(start, end int, cost [][]int) int {
	if start == end {
		return 0
	}
	if cost[start][end] != 0 {
		return cost[start][end]
	}
	cost[start][end] = 20100
	for i := start + 1; i < end; i++ {
		left := dp(start, i-1, cost)
		right := dp(i+1, end, cost)
		if left < right {
			left = right
		}
		if c := i + left; cost[start][end] > c {
			cost[start][end] = c
		}
	}
	return cost[start][end]
}

func getMoneyAmount(n int) int {
	// 1 <= n <= 200
	cost := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		cost[i] = make([]int, n+1)
	}
	for i := 1; i < n; i++ {
		cost[i][i+1] = i
	}
	return dp(1, n, cost)
}
