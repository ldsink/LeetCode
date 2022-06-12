package main

func findAndReplacePattern(words []string, pattern string) []string {
	var result []string
	for _, word := range words {
		if len(word) != len(pattern) {
			continue
		}
		p, q := make([]rune, 26), make([]rune, 26)
		match := true
		for i, r := range []rune(word) {
			a := int(r - 'a')
			b := int(pattern[i] - 'a')
			if p[a] == 0 && q[b] == 0 {
				p[a] = rune(pattern[i])
				q[b] = r
				continue
			} else if p[a] != rune(pattern[i]) || q[b] != r {
				match = false
				break
			}
		}
		if match {
			result = append(result, word)
		}
	}
	return result
}
