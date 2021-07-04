package main

func countBinarySubstrings(s string) int {
	var result, prevCount int
	r := []rune(s)
	for ; len(r) > 0; {
		var end int
		for end = 1; end < len(r); end++ {
			if r[end] != r[0] {
				break
			}
		}
		if end > prevCount {
			result += prevCount
		} else {
			result += end
		}
		prevCount = end
		r = r[end:]
	}
	return result
}
