package main

import (
	"sort"
)

func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	var result []bool
	for i, left := range l {
		right := r[i]
		if left == right {
			result = append(result, true)
			continue
		}
		current := append([]int{}, nums[left:right+1]...)
		sort.Ints(current)
		if (current[len(current)-1]-current[0])%(len(current)-1) != 0 {
			result = append(result, false)
			continue
		}
		step := (current[len(current)-1] - current[0]) / (len(current) - 1)
		match := true
		for j := 0; j < len(current)-1; j++ {
			if current[j]+step != current[j+1] {
				match = false
				break
			}
		}
		result = append(result, match)
	}
	return result
}
