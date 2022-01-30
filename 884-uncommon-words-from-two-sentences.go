package main

import "strings"

func uncommonFromSentences(s1 string, s2 string) []string {
	count := make(map[string]int)
	for _, word := range strings.Split(s1, " ") {
		count[word]++
	}
	for _, word := range strings.Split(s2, " ") {
		count[word]++
	}
	var result []string
	for word, c := range count {
		if c == 1 {
			result = append(result, word)
		}
	}
	return result
}
