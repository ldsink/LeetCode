package main

import (
	"sort"
)

func _enumSubSets(nums []int, used []bool, index int) [][]int {
	if index == len(nums) {
		var result []int
		for i := 0; i < len(nums); i++ {
			if used[i] {
				result = append(result, nums[i])
			}
		}
		if len(result) > 0 {
			return [][]int{result}
		} else {
			return [][]int{}
		}
	}

	var result [][]int
	used[index] = true
	r := _enumSubSets(nums, used, index+1)
	result = append(result, r...)
	used[index] = false
	r = _enumSubSets(nums, used, index+1)
	result = append(result, r...)
	return result
}

func subsets(nums []int) [][]int {
	result := [][]int{[]int{}}
	used := make([]bool, len(nums))
	r := _enumSubSets(nums, used, 0)
	result = append(result, r...)
	return result
}

func singleNumSubSets(num, max int) [][]int {
	var result [][]int
	var current []int
	for i := 1; i <= max; i++ {
		c := append(current, num)
		result = append(result, c)
		current = c
	}
	return result
}

func extendSubSets(nums []int, numCount map[int]int) [][]int {
	num := nums[0]
	curr := singleNumSubSets(num, numCount[num])
	if len(nums) == 1 {
		return singleNumSubSets(num, numCount[num])
	}
	var result [][]int
	rest := extendSubSets(nums[1:], numCount)
	for _, c := range curr {
		for _, e := range rest {
			n := append(c, e...)
			result = append(result, n)
		}
	}
	return result
}

func subsetsWithDup(nums []int) [][]int {
	numCount := make(map[int]int)
	for _, num := range nums {
		numCount[num] += 1
	}

	uniqueNums := make([]int, len(numCount))
	i := 0
	for num := range numCount {
		uniqueNums[i] = num
		i++
	}
	sort.Ints(uniqueNums)
	uniqueSubSets := subsets(uniqueNums)

	result := [][]int{[]int{}}
	for _, sets := range uniqueSubSets {
		if len(sets) == 0 {
			continue
		}
		extend := extendSubSets(sets, numCount)
		for _, r := range extend {
			result = append(result, r)
		}
	}

	return result
}
