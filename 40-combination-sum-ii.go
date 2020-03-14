package main

import (
	"sort"
	"strconv"
	"strings"
)

func getSliceString(s []int) string {
	var nums []string
	for _, val := range s {
		text := strconv.Itoa(val)
		nums = append(nums, text)
	}
	return strings.Join(nums, ",")
}

func checkSortedCandidates2(candidates []int, index, target int) [][]int {
	var result [][]int
	for i := index; i >= 0; i-- {
		if target < candidates[i] {
			continue
		}
		if target == candidates[i] {
			result = append(result, []int{target})
			continue
		}
		r := checkSortedCandidates2(candidates, i-1, target-candidates[i])
		if len(r) == 0 {
			continue
		}
		for _, line := range r {
			result = append(result, append(line, candidates[i]))
		}
	}
	return result
}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var result [][]int
	m := make(map[string]bool)
	for _, r := range checkSortedCandidates2(candidates, len(candidates)-1, target) {
		h := getSliceString(r)
		if m[h] {
			continue
		}
		m[h] = true
		result = append(result, r)
	}
	return result
}
