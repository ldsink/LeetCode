package main

import (
	"sort"
)

func checkSortedCandidates(candidates []int, index, target int) [][]int {
	var result [][]int
	for i := index; i >= 0; i-- {
		if target < candidates[i] {
			continue
		}
		if target == candidates[i] {
			result = append(result, []int{target})
			continue
		}
		var s []int
		for j, k := 1, target/candidates[i]; j <= k; j++ {
			s = append(s, candidates[i])
			rest := target - (candidates[i] * j)
			if rest == 0 {
				result = append(result, s)
			}
			r := checkSortedCandidates(candidates, i-1, rest)
			if len(r) == 0 {
				continue
			}
			for _, line := range r {
				result = append(result, append(line, s...))
			}
		}
	}
	return result
}

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	return checkSortedCandidates(candidates, len(candidates)-1, target)
}
