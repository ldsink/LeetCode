package main

import "strings"

func lengthOfLastWord(s string) int {
	words := strings.Split(s, " ")
	for i := len(words) - 1; i >= 0; i-- {
		if len(words[i]) > 0 {
			return len(words[i])
		}
	}
	return 0
}
