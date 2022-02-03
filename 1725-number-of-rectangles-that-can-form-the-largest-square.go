package main

func countGoodRectangles(rectangles [][]int) int {
	var maxLen int
	count := make(map[int]int)
	for _, data := range rectangles {
		a, b := data[0], data[1]
		if a > b {
			a = b
		}
		count[a]++
		if maxLen < a {
			maxLen = a
		}
	}
	return count[maxLen]
}
