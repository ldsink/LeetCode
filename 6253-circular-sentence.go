package main

import "strings"

func isCircularSentence(sentence string) bool {
	words := strings.Split(sentence, " ")
	for i := 0; i < len(words); i++ {
		if words[i][len(words[i])-1] != words[(i+1)%len(words)][0] {
			return false
		}
	}
	return true
}
