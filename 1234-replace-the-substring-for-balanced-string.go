package main

func balancedString(s string) int {
	var delta [4]int
	mapping := make(map[rune]int)
	mapping['Q'] = 0
	mapping['W'] = 1
	mapping['E'] = 2
	mapping['R'] = 3
	for _, r := range s {
		delta[mapping[r]]++
	}
	avg := len(s) >> 2
	var minCount int
	for i := 0; i < 4; i++ {
		if delta[i] >= avg {
			delta[i] -= avg
		} else {
			delta[i] = 0
		}
		minCount += delta[i]
	}
	result := len(s)
	var current [4]int
	canBalance := func() bool {
		for i := 0; i < 4; i++ {
			if current[i] < delta[i] {
				return false
			}
		}
		return true
	}
	for start, end := 0, 0; start < len(s); start++ {
		if start > 0 {
			current[mapping[rune(s[start-1])]]--
		}
		for ; end < len(s) && !canBalance(); end++ {
			current[mapping[rune(s[end])]]++
		}
		if result > end-start && canBalance() {
			result = end - start
		}
		if result == minCount { // 已到达最优解，提前退出
			return result
		}
	}
	return result
}
