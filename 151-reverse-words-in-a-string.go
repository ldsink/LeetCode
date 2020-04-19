package main

import (
	"strings"
)

func reverseWords(s string) string {
	ss := strings.Split(strings.TrimSpace(s), " ")
	var words []string
	for i := len(ss) - 1; i >= 0; i-- {
		if ss[i] == "" {
			continue
		}
		words = append(words, ss[i])
	}
	return strings.Join(words, " ")
}
