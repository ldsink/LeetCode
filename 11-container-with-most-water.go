package main

import "fmt"

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func maxArea(height []int) int {
	heightBefore := make(map[int]int)
	heightAfter := make(map[int]int)
	submerged := make(map[int]bool)

	for i, j := 1, height[0]; i < len(height); i++ {
		heightBefore[i] = j
		if j < height[i] {
			j = height[i]
		}
	}
	for i, j := len(height)-2, height[len(height)-1]; i >= 0; i-- {
		heightAfter[i] = j
		if j < height[i] {
			j = height[i]
		}
	}
	for i, j := 1, len(height)-2; i <= j; i++ {
		if heightBefore[i] >= height[i] && height[i] <= heightAfter[i] {
			submerged[i] = true
		}
	}

	result := 0
	for i := 0; i < len(height)-1; i++ {
		if submerged[i] {
			continue
		}
		for j := i + 1; j < len(height); j++ {
			if submerged[j] {
				continue
			}
			area := Min(height[i], height[j]) * (j - i)
			if result < area {
				result = area
			}
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
