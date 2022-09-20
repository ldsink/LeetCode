package main

import (
	"sort"
)

func canPartitionKSubsets(nums []int, k int) bool {
	var sum int
	for _, num := range nums {
		sum += num
	}
	if sum%k != 0 {
		return false
	}
	target := sum / k
	for _, num := range nums {
		if num > target {
			return false
		}
	}
	sort.Slice(nums, func(i, j int) bool { // desc order
		return nums[i] > nums[j]
	})

	buckets := make([]int, k)
	var search func(int) bool
	search = func(idx int) bool {
		if idx == len(nums) {
			return true
		}
		for i := 0; i < k; i++ {
			if i > 0 && buckets[i-1] == buckets[i] { // skip duplicate conditions
				continue
			}
			if nums[idx]+buckets[i] <= target {
				buckets[i] += nums[idx]
				if search(idx + 1) {
					return true
				}
				buckets[i] -= nums[idx]
			}
		}
		return false
	}
	return search(0)
}
