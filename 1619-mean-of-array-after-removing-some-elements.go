package main

import "sort"

func trimMean(arr []int) float64 {
	sort.Ints(arr)
	l := len(arr) / 20
	arr = arr[l : len(arr)-l]
	var sum int
	for _, num := range arr {
		sum += num
	}
	return float64(sum) / float64(len(arr))
}
