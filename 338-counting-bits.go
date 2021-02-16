package main

import (
	"math/bits"
)

func countBits(num int) []int {
	if num == 0 {
		return []int{0}
	}
	result := []int{0, 1}
	if num == 1 {
		return result
	}
	for i := uint(2); i <= uint(num); i++ {
		offset := i - (1 << (bits.Len(i) - 1))
		result = append(result, 1+result[offset])
	}
	return result
}
