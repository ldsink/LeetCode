package main

import "strings"

func halvesAreAlike(s string) bool {
	vowels := "aeoiuAEIOU"
	var a, b int
	for i, j := 0, len(s)/2; i < j; i++ {
		if strings.Contains(vowels, s[i:i+1]) {
			a++
		}
	}
	for i := len(s) / 2; i < len(s); i++ {
		if strings.Contains(vowels, s[i:i+1]) {
			b++
		}
	}
	return a == b
}
