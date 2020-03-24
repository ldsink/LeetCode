package main

func largestRectangleArea(heights []int) int {
	leftHigh := make([]int, len(heights))
	rightHigh := make([]int, len(heights))
	for i := 0; i < len(heights); i++ {
		for j := i - 1; j >= 0; j-- {
			if heights[j] >= heights[i] {
				leftHigh[i] += leftHigh[j] + 1
				j -= leftHigh[j]
				continue
			}
			break
		}
	}
	for i := len(heights) - 1; i >= 0; i-- {
		for j := i + 1; j < len(heights); j++ {
			if heights[j] >= heights[i] {
				rightHigh[i] += rightHigh[j] + 1
				j += rightHigh[j]
				continue
			}
			break
		}
	}
	var result int
	for i := 0; i < len(heights); i++ {
		area := (1 + leftHigh[i] + rightHigh[i]) * heights[i]
		if result < area {
			result = area
		}
	}
	return result
}
