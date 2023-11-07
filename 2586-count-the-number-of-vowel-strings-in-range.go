package main

import "strings"

func vowelStrings(words []string, left int, right int) int {
	var result int
	vowels := "aeoiuAEIOU"
	for ; left <= right; left++ {
		if strings.ContainsAny(vowels, words[left][0:1]) && strings.ContainsAny(vowels, words[left][len(words[left])-1:]) {
			result++
		}
	}
	return result
}
