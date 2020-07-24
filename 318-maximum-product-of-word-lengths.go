package main

import (
	"sort"
)

func maxProduct(words []string) int {
	var result int
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) > len(words[j])
	})

	letters := make([][26]bool, len(words))
	for i := 0; i < len(words); i++ {
		for _, c := range words[i] {
			letters[i][c-'a'] = true
		}
	}

	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if result > len(words[i])*len(words[j]) {
				break
			}
			common := false
			for k := 0; k < 26; k++ {
				if letters[i][k] && letters[j][k] {
					common = true
					break
				}
			}
			if !common {
				result = len(words[i]) * len(words[j])
				break
			}
		}
	}
	return result
}
