package main

import (
	"math"
)

func minOperations(nums []int, x int) int {
	if nums[0] == x || nums[len(nums)-1] == x {
		return 1
	}
	rightSum := make(map[int]int)
	for i, j := 1, 0; i <= len(nums); i++ {
		j += nums[len(nums)-i]
		rightSum[j] = i
	}
	result := math.MaxInt
	for i, j := 0, 0; i < len(nums); i++ {
		if j == x {
			if result > i {
				result = i
			}
		} else if c, ok := rightSum[x-j]; ok {
			if result > i+c {
				result = i + c
			}
		}
		j += nums[i]
	}
	if result > len(nums) {
		return -1
	}
	return result
}
