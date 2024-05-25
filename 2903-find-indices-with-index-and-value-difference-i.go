package main

func findIndices(nums []int, indexDifference int, valueDifference int) []int {
	abs := func(x int) int {
		if x >= 0 {
			return x
		}
		return -x
	}
	for i := 0; i+indexDifference < len(nums); i++ {
		for j := i + indexDifference; j < len(nums); j++ {
			if abs(nums[j]-nums[i]) >= valueDifference {
				return []int{i, j}
			}
		}
	}
	return []int{-1, -1}
}
