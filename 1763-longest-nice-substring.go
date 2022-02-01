package main

func longestNiceSubstring(s string) string {
	if len(s) < 2 {
		return ""
	}
	lower := make([]bool, 26)
	upper := make([]bool, 26)
	for _, r := range []rune(s) {
		if 'a' <= r && r <= 'z' {
			lower[r-'a'] = true
		} else {
			upper[r-'A'] = true
		}
	}
	var subS []string
	var i, j int
	for i, j = 0, 0; i < len(s); i++ {
		var cIdx int
		if 'a' <= s[i] && s[i] <= 'z' {
			cIdx = int(s[i] - 'a')
		} else {
			cIdx = int(s[i] - 'A')
		}
		if lower[cIdx] && upper[cIdx] {
			continue
		}
		if j < i {
			subS = append(subS, s[j:i])
		}
		j = i + 1
	}
	if j == 0 {
		return s
	}
	if j < i {
		subS = append(subS, s[j:i])
	}
	var result string
	for _, s := range subS {
		c := longestNiceSubstring(s)
		if len(result) < len(c) {
			result = c
		}
	}
	return result
}
