package main

import (
	"math"
	"sort"
)

func makeArrayIncreasing(arr1 []int, arr2 []int) int {
	if len(arr1) == 0 {
		return 0
	}
	sort.Ints(arr2)
	findGte := func(val int) int {
		start, end := 0, len(arr2)-1
		for start != end {
			middle := (start + end) >> 1
			if val >= arr2[middle] {
				start = middle + 1
			} else {
				end = middle
			}
		}
		if arr2[start] > val {
			return start
		}
		return -1
	}
	result := make(map[int]int)
	result[arr1[0]] = 0
	if arr2[0] < arr1[0] {
		result[arr2[0]] = 1
	}
	for i := 1; i < len(arr1) && len(result) > 0; i++ {
		temp := make(map[int]int)
		for val, count := range result {
			if val < arr1[i] {
				if c, ok := temp[arr1[i]]; !ok || c > count {
					temp[arr1[i]] = count
				}
			}
			if idx := findGte(val); idx != -1 {
				if c, ok := temp[arr2[idx]]; !ok || c > count+1 {
					temp[arr2[idx]] = count + 1
				}
			}
		}
		result = temp
	}
	if len(result) == 0 {
		return -1
	}
	count := math.MaxInt
	for _, c := range result {
		if count > c {
			count = c
		}
	}
	return count
}
