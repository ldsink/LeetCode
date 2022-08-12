package main

import "sort"

func maxChunksToSorted(arr []int) int {
	sorted := make([]int, len(arr))
	for idx, num := range arr {
		sorted[idx] = num
	}
	sort.Ints(sorted)

	hashA, hashB := make(map[int]int), make(map[int]int)
	isEqual := func(a, b map[int]int) bool {
		if len(a) != len(b) {
			return false
		}
		for k := range a {
			if a[k] != b[k] {
				return false
			}
		}
		return true
	}

	var result int
	for i := 0; i < len(arr); i++ {
		hashA[arr[i]]++
		hashB[sorted[i]]++
		if isEqual(hashA, hashB) {
			result++
			hashA, hashB = make(map[int]int), make(map[int]int)
		}
	}
	return result
}
