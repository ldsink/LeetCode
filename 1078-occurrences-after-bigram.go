package main

import "strings"

func findOcurrences(text string, first string, second string) []string {
	var result []string
	words := strings.Split(text, " ")
	for i := 0; i < len(words)-2; i++ {
		if words[i] == first && words[i+1] == second {
			result = append(result, words[i+2])
		}
	}
	return result
}
