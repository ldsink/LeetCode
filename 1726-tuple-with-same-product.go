package main

import (
	"sort"
)

func tupleSameProduct(nums []int) int {
	sort.Ints(nums)
	sum := make(map[int]int)
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			sum[nums[i]*nums[j]]++
		}
	}
	// C(N, 2) * 8
	var result int
	for _, n := range sum {
		if n >= 2 {
			result += 4 * n * (n - 1) // 8 * (n * (n-1)) / (2 * 1)
		}
	}
	return result
}
