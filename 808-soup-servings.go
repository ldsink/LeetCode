package main

import "math"

func soupServings(n int) float64 {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	n = int(math.Ceil(float64(n) / 25))
	if n >= 179 {
		return 1
	}
	rate := make([][]float64, n+1)
	for i := range rate {
		rate[i] = make([]float64, n+1)
	}
	rate[0][0] = 0.5
	for i := 1; i <= n; i++ {
		rate[0][i] = 1
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			k := rate[max(0, i-4)][j] + rate[max(0, i-3)][max(0, j-1)] + rate[max(0, i-2)][max(0, j-2)] + rate[max(0, i-1)][max(0, j-3)]
			rate[i][j] = k / 4
		}
	}
	return rate[n][n]
}
