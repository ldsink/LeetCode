package main

import (
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	mark := make(map[int]map[int]map[int]bool)
	var result [][]int
	for i := 0; i < len(nums)-3; i++ {
		for j := len(nums) - 1; j > i+2; j-- {
			for k := i + 1; k < j-1; k++ {
				if _, ok := mark[nums[i]]; ok {
					if _, ok := mark[nums[i]][nums[j]]; ok {
						if _, ok := mark[nums[i]][nums[j]][nums[k]]; ok {
							continue
						}
					}
				}

				l := target - nums[i] - nums[j] - nums[k]
				if !(nums[k] <= l && l <= nums[j]) {
					continue
				}

				success := false
				for left, right := k+1, j-1; left <= right; {
					if left == right {
						if nums[left] == l {
							success = true
						}
						break
					}
					middle := (left + right) >> 1
					if l <= nums[middle] {
						right = middle
					} else {
						left = middle + 1
					}
				}

				if success {
					result = append(result, []int{nums[i], nums[k], l, nums[j]})
				}

				if _, ok := mark[nums[i]]; !ok {
					mark[nums[i]] = make(map[int]map[int]bool)
				}
				if _, ok := mark[nums[i]][nums[j]]; !ok {
					mark[nums[i]][nums[j]] = make(map[int]bool)
				}
				mark[nums[i]][nums[j]][nums[k]] = success
			}
		}
	}
	return result
}
