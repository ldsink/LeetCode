package main

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func maxArea(height []int) int {
	var result int
	for i, j := 0, len(height)-1; i < j; {
		area := Min(height[i], height[j]) * (j - i)
		if result < area {
			result = area
		}
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return result
}
