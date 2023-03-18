package main

import (
	"sort"
)

func answerQueries(nums []int, queries []int) []int {
	sort.Ints(nums)
	sum := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		sum[i+1] = sum[i] + nums[i]
	}
	var start, end int
	var result []int
	for _, num := range queries {
		if sum[len(nums)] <= num {
			result = append(result, len(nums))
		} else if num < sum[1] {
			result = append(result, 0)
		} else {
			for start, end = 0, len(nums); start != end; {
				middle := (start + end) >> 1
				if num >= sum[middle+1] {
					start = middle + 1
				} else {
					end = middle
				}
			}
			result = append(result, start)
		}
	}
	return result
}
