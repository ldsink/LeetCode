package main

import "fmt"

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

func main() {
	var height []int
	height = []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(height))
	height = []int{3, 6}
	fmt.Println(maxArea(height))
}
