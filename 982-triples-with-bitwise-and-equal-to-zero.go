package main

import "sort"

func countTriplets(nums []int) int {
	var result int
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			for k := j; k < len(nums); k++ {
				if nums[i]&nums[j]&nums[k] == 0 {
					if i == j && j == k {
						result += 1
					} else if i == j || j == k {
						result += 3
					} else {
						result += 6
					}
				}
			}
		}
	}
	return result
}
