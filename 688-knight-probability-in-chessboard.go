package main

func dp(n, k, row, column int, directions [8][2]int, cached map[[3]int]float64) float64 {
	if value, ok := cached[[3]int{row, column, k}]; ok {
		return value
	}
	if row < 0 || row >= n || column < 0 || column >= n {
		cached[[3]int{row, column, k}] = 0
		return 0
	}
	if k == 0 {
		cached[[3]int{row, column, k}] = 1
		return 1
	}
	var rate float64
	for _, direction := range directions {
		rate += 0.125 * dp(n, k-1, row+direction[0], column+direction[1], directions, cached)
	}
	cached[[3]int{row, column, k}] = rate
	return rate
}

func knightProbability(n int, k int, row int, column int) float64 {
	directions := [8][2]int{{-2, -1}, {-2, 1}, {-1, 2}, {1, 2}, {2, 1}, {2, -1}, {1, -2}, {-1, -2}}
	cached := make(map[[3]int]float64)
	return dp(n, k, row, column, directions, cached)
}
