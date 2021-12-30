package main

import (
	"sort"
)

func canArrange(arr []int, k int) bool {
	for i := 0; i < len(arr); i++ {
		arr[i] %= k
	}
	var nonZero []int
	var zero int
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			zero++
		} else if arr[i] > 0 {
			nonZero = append(nonZero, arr[i])
		} else {
			nonZero = append(nonZero, k+arr[i])
		}
	}
	if zero%2 != 0 {
		return false
	}
	sort.Ints(nonZero)
	for i, j := 0, len(nonZero)-1; i < len(nonZero)/2; i++ {
		if (nonZero[i]+nonZero[j-i])%k != 0 {
			return false
		}
	}
	return true
}
