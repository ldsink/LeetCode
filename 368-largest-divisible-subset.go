package main

import "sort"

func largestDivisibleSubset(nums []int) []int {
	sort.Ints(nums)
	var subset [][]int
	var result []int
	for i := 0; i < len(nums); i++ {
		subset = append(subset, []int{nums[i]})
		for j := i - 1; j >= 0; j-- {
			if len(subset[i]) > len(subset[j]) {
				continue
			}
			satisfied := true
			for _, k := range subset[j] {
				if k%nums[i] == 0 || nums[i]%k == 0 {
					continue
				}
				satisfied = false
				break
			}
			if satisfied {
				subset[i] = append([]int{nums[i]}, subset[j]...)
			}
		}
		if len(result) < len(subset[i]) {
			result = subset[i]
		}
	}
	return result
}
