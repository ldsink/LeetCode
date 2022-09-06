package main

import (
	"strings"
)

func reorderSpaces(text string) string {
	var words []string
	for _, word := range strings.Split(text, " ") {
		if word != "" {
			words = append(words, word)
		}
	}
	count := len(text) - len(strings.Join(words, ""))
	singleCount := 0
	if len(words) > 1 {
		singleCount = count / (len(words) - 1)
	}
	restCount := count - singleCount*(len(words)-1)

	getSpaceString := func(n int) string {
		var elems []string
		for i := 0; i < n; i++ {
			elems = append(elems, " ")
		}
		return strings.Join(elems, "")
	}

	return strings.Join(words, getSpaceString(singleCount)) + getSpaceString(restCount)
}
