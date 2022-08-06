package main

import (
	"sort"
	"strings"
)

func stringMatching(words []string) []string {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
	var result []string
	for i := 0; i < len(words)-1; i++ {
		for j := i + 1; j < len(words); j++ {
			if len(words[j]) == len(words[i]) {
				continue
			}
			if strings.Contains(words[j], words[i]) {
				result = append(result, words[i])
				break
			}
		}
	}
	return result
}
