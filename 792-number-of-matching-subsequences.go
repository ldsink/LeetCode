package main

func numMatchingSubseq(s string, words []string) int {
	next := make([]map[int]int, len(s))
	var count int
	isSubseq := func(word string) bool {
		i, j, k := 0, 0, 0
		for i < len(s) && j < len(word) {
			if next[i] == nil {
				next[i] = make(map[int]int)
			}
			if k, ok := next[i][int(word[j])]; ok {
				if k < len(s) {
					i = k + 1
					j++
					continue
				}
				return false
			}
			for k = i; k < len(s); k++ {
				if s[k] == word[j] {
					break
				} else if _, ok := next[i][int(s[k])]; !ok {
					next[i][int(s[k])] = k
				}
			}
			next[i][int(word[j])] = k
			if k == len(s) {
				return false
			} else {
				i = k + 1
				j++
			}
		}
		return j == len(word)
	}
	for _, word := range words {
		if isSubseq(word) {
			count++
		}
	}
	return count
}
