package main

import "strings"

func maxRepeating(sequence string, word string) int {
	var curr string
	for i, j := 1, len(sequence)/len(word); i <= j; i++ {
		curr += word
		if !strings.Contains(sequence, curr) {
			return i - 1
		}
	}
	return len(sequence) / len(word)
}
