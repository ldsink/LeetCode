package main

import "sort"

func largestSumAfterKNegations(nums []int, k int) int {
	if k == 0 || len(nums) == 0 {
		var result int
		for _, num := range nums {
			result += num
		}
		return result
	}
	sort.Ints(nums)
	var minusCount int
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			minusCount++
		} else {
			break
		}
	}

	if minusCount == 0 {
		if k%2 != 0 {
			nums[0] = -nums[0]
		}
		return largestSumAfterKNegations(nums, 0)
	}

	if minusCount <= k {
		for i := 0; i < minusCount; i++ {
			nums[i] = -nums[i]
		}
		return largestSumAfterKNegations(nums, k-minusCount)
	}
	for i := 0; i < k; i++ {
		nums[i] = -nums[i]
	}
	return largestSumAfterKNegations(nums, 0)
}
