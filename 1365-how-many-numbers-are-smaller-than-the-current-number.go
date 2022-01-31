package main

import "sort"

func smallerNumbersThanCurrent(nums []int) []int {
	count := make(map[int]int)
	for _, n := range nums {
		count[n]++
	}

	var array []int
	for n, _ := range count {
		array = append(array, n)
	}
	sort.Ints(array)
	smaller := make(map[int]int)
	var acc int
	for _, n := range array {
		smaller[n] = acc
		acc += count[n]
	}

	var result []int
	for _, n := range nums {
		result = append(result, smaller[n])
	}
	return result
}
