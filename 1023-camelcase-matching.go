package main

func camelMatch(queries []string, pattern string) []bool {
	var result []bool
	countUpper := func(query string) (c int) {
		for _, r := range query {
			if 'A' <= r && r <= 'Z' {
				c++
			}
		}
		return
	}
	pCount := countUpper(pattern)
	isMatch := func(query string) bool {
		if countUpper(query) != pCount {
			return false
		}
		var i int
		for _, r := range query {
			if r == rune(pattern[i]) {
				i++
			}
			if i == len(pattern) {
				break
			}
		}
		return i == len(pattern)
	}
	for _, query := range queries {
		result = append(result, isMatch(query))
	}
	return result
}
