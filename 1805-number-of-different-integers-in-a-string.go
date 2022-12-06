package main

import "strings"

func numDifferentIntegers(word string) int {
	parseInt := func(idx int) (string, int) {
		var start, end int
		for start = idx; start < len(word); start++ {
			if '0' <= word[start] && word[start] <= '9' {
				break
			}
		}
		if start == len(word) {
			return "", -1
		}
		for end = start + 1; end < len(word); end++ {
			if '0' <= word[end] && word[end] <= '9' {
				continue
			}
			break
		}
		return strings.TrimLeft(word[start:end], "0"), end
	}
	nums := make(map[string]bool)
	for idx, val := 0, ""; idx != -1; {
		val, idx = parseInt(idx)
		if idx != -1 {
			nums[val] = true
		}
	}
	return len(nums)
}
