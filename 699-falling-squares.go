package main

func fallingSquares(positions [][]int) []int {
	n := len(positions)
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}
	heights := make([]int, n)
	for i, position := range positions {
		currLeft, currRight := position[0], position[0]+position[1]-1
		heights[i] = position[1]
		for j, prevPosition := range positions[:i] {
			prevLeft, prevRight := prevPosition[0], prevPosition[0]+prevPosition[1]-1
			if prevLeft <= currRight && currLeft <= prevRight {
				heights[i] = max(heights[i], heights[j]+position[1])
			}
		}
	}
	for i := 1; i < n; i++ {
		heights[i] = max(heights[i], heights[i-1])
	}
	return heights
}
