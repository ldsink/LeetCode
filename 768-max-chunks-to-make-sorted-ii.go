package main

import (
	"sort"
)

func maxChunksToSorted(arr []int) int {
	sorted := make([]int, len(arr))
	for idx, num := range arr {
		sorted[idx] = num
	}
	sort.Ints(sorted)

	var result, sumA, sumB, xorA, xorB int
	for idx, num := range sorted {
		sumA += arr[idx]
		xorA ^= arr[idx]
		sumB += num
		xorB ^= num
		// 用 sum 和 xor 检查，防止越界造成的错误 sum 相等
		if sumA == sumB && xorA == xorB {
			result++
		}
	}
	return result
}
