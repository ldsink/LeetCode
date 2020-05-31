package main

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func numSquares(n int) int {
	if n == 0 {
		return 0
	}

	maxInt := int(^uint(0) >> 1)
	dp := make([]int, n+1)
	for i := 1; i < len(dp); i++ {
		dp[i] = maxInt
	}
	for i := 0; i <= n; i++ {
		for j := 1; i+j*j <= n; j++ {
			dp[i+j*j] = min(dp[i+j*j], dp[i]+1)
		}
	}
	return dp[n]
}
