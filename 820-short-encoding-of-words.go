package main

import (
	"sort"
	"strings"
)

func minimumLengthEncoding(words []string) int {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) > len(words[j])
	})
	var result int
	used := make([]bool, len(words))
	for i := 0; i < len(words); i++ {
		if used[i] {
			continue
		}
		used[i] = true
		result += len(words[i]) + 1
		for j := i + 1; j < len(words); j++ {
			if used[j] {
				continue
			}
			if strings.HasSuffix(words[i], words[j]) {
				used[j] = true
			}
		}
	}
	return result
}
