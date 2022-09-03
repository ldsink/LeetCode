package main

import (
	"sort"
)

func findLongestChain(pairs [][]int) int {
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][0] < pairs[j][0] || (pairs[i][0] == pairs[j][0] && pairs[i][1] < pairs[j][1])
	})
	count := make([]int, len(pairs))
	var result int
	for i := len(pairs) - 1; i >= 0; i-- {
		for j := len(pairs) - 1; j > i; j-- {
			if pairs[j][0] <= pairs[i][1] {
				break
			}
			if count[i] < count[j] {
				count[i] = count[j]
				if count[i] == result { // 已经是之前的最优解，停止循环
					break
				}
			}
		}
		count[i]++
		if result < count[i] {
			result = count[i]
		}
	}
	return result
}
