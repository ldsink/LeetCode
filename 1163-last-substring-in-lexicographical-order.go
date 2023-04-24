package main

import "sort"

func lastSubstring(s string) string {
	var indexes []int
	maxR := 'a' - 1
	for i, r := range s {
		if maxR < r {
			maxR, indexes = r, []int{i}
		} else if maxR == r {
			indexes = append(indexes, i)
		}
	}
	// 当前的 r 一定是最大的，从后往前，如果相邻的，后者肯定比前者小
	temp := []int{indexes[0]}
	for i := len(indexes) - 1; i > 0; i-- {
		if indexes[i-1]+1 == indexes[i] {
			continue
		} else {
			temp = append(temp, indexes[i])
		}
	}
	indexes = temp
	for offset, same := 1, 0; len(indexes) != 1 && same <= 8; offset++ {
		origin := len(indexes)
		candidates := [26][]int{}
		for _, i := range indexes {
			if i+offset < len(s) {
				r := int(s[i+offset] - 'a')
				candidates[r] = append(candidates[r], i)
			}
		}
		for r := maxR - 'a'; r >= 0; r-- {
			if len(candidates[r]) > 0 {
				indexes = candidates[r]
				break
			}
		}
		if origin-len(indexes) <= 1 {
			same++
		} else {
			same = 0
		}
	}
	sort.Ints(indexes)
	return s[indexes[0]:]
}
