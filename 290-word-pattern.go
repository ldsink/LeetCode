package main

import (
	"strings"
)

func wordPattern(pattern string, str string) bool {
	r2w := make(map[byte]string)
	w2r := make(map[string]byte)
	words := strings.Split(str, " ")
	if len(pattern) != len(words) {
		return false
	}
	for i, word := range words {
		if w, ok := r2w[pattern[i]]; ok {
			if w != word {
				return false
			}
		} else {
			r2w[pattern[i]] = word
		}

		if r, ok := w2r[word]; ok {
			if r != pattern[i] {
				return false
			}
		} else {
			w2r[word] = pattern[i]
		}
	}
	return true
}
