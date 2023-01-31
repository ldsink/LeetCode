package main

import (
	"unicode"
)

func decodeMessage(key string, message string) string {
	dict := make(map[rune]rune)
	base := 'a'
	for _, r := range key {
		if _, ok := dict[r]; unicode.IsLower(r) && !ok {
			dict[r] = base
			base++
		}
	}
	var result []rune
	for _, r := range message {
		if nr, ok := dict[r]; ok {
			result = append(result, nr)
		} else {
			result = append(result, r)
		}
	}
	return string(result)
}
