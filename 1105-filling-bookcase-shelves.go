package main

func minHeightShelves(books [][]int, shelfWidth int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	dp := make([]int, len(books)+1)
	for i := 1; i <= len(books); i++ {
		dp[i] = 1000000
	}
	for i := 0; i < len(books); i++ {
		maxHeight, curWidth := 0, 0
		for j := i; j >= 0; j-- {
			curWidth += books[j][0]
			if curWidth > shelfWidth {
				break
			}
			maxHeight = max(maxHeight, books[j][1])
			dp[i+1] = min(dp[i+1], dp[j]+maxHeight)
		}
	}
	return dp[len(books)]
}
