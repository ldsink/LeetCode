package main

import (
	"regexp"
	"strings"
)

func countValidWords(sentence string) int {
	parts := strings.Split(sentence, " ")
	var count int
	regexp1, _ := regexp.Compile("^[a-z]*[!.,]?$")
	regexp2, _ := regexp.Compile("^[a-z]+-[a-z]+[!.,]?$")
	for _, s := range parts {
		if len(s) == 0 {
			continue
		}
		if regexp1.MatchString(s) || regexp2.MatchString(s) {
			count++
		}
	}
	return count
}
