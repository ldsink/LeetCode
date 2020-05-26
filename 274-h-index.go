package main

import "sort"

func hIndex(citations []int) int {
	sort.Ints(citations)
	for i, j := len(citations)-1, 1; i >= 0; i, j = i-1, j+1 {
		if citations[i] < j {
			return j - 1
		}
	}
	return len(citations)
}
