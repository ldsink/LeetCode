package main

import "math"

func findClosest(words []string, word1 string, word2 string) int {
	word1Idx := -1
	word2Idx := -1
	result := math.MaxInt
	for idx, word := range words {
		if word == word1 {
			if word2Idx != -1 && result > idx-word2Idx {
				result = idx - word2Idx
			}
			word1Idx = idx
		} else if word == word2 {
			if word1Idx != -1 && result > idx-word1Idx {
				result = idx - word1Idx
			}
			word2Idx = idx
		}
	}
	return result
}
