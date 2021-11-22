package main

func buddyStrings(s string, goal string) bool {
	if len(s) != len(goal) || len(s) < 2 {
		return false
	}
	// 如果两个字符串相同，那有一个字母重复两次，就 OK
	if s == goal {
		count := make([]int, 26)
		for _, r := range []rune(s) {
			count[r-'a']++
			if count[r-'a'] > 1 {
				return true
			}
		}
		return false
	}
	// 否则两个字符串差别只能有两个字符
	var diffIdx []int
	for i := 0; i < len(s); i++ {
		if s[i] != goal[i] {
			diffIdx = append(diffIdx, i)
			if len(diffIdx) > 2 {
				return false
			}
		}
	}
	if len(diffIdx) != 2 {
		return false
	}
	return s[diffIdx[0]] == goal[diffIdx[1]] && s[diffIdx[1]] == goal[diffIdx[0]]
}
