package main

import "regexp"

func longestSubstring(s string, k int) int {
	if len(s) < k {
		return 0
	}

	// 统计每个字符的数量
	count := make(map[rune]int)
	for _, c := range s {
		count[c]++
	}

	// 取出不满足数量的字符
	invalid := make(map[rune]bool)
	for r, c := range count {
		if c < k {
			invalid[r] = true
		}
	}

	// 如果都满足，那么当前串就满足条件
	if len(invalid) == 0 {
		return len(s)
	}

	// 否则按照不满足的字符拆分字符串，递归求解
	var invalidRunes []rune
	for r, _ := range invalid {
		invalidRunes = append(invalidRunes, r)
	}
	re := regexp.MustCompile("[" + string(invalidRunes) + "]")
	var result int
	for _, part := range re.Split(s, -1) {
		if len(part) == 0 {
			continue
		}
		if r := longestSubstring(part, k); result < r {
			result = r
		}
	}
	return result
}
