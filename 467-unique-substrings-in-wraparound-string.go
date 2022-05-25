package main

import (
	"strings"
)

func findSubstringInWraproundString(p string) int {
	const base = "abcdefghijklmnopqrstuvwxyz"
	current := base
	var count int
	for i, prev := 0, true; i < len(p) && prev; i++ {
		if i+len(base) > len(current) {
			current += base
		}
		prev = false
		for j := 0; j < len(base); j++ {
			if strings.Contains(p, current[j:j+i+1]) {
				count++
				prev = true
			}
		}
	}
	return count
}
