package main

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

func countKDifference(nums []int, k int) int {
	var count int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if abs(nums[j]-nums[i]) == k {
				count++
			}
		}
	}
	return count
}
