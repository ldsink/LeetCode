package main

import (
	"strings"
)

func getSpace(n int) string {
	s := ""
	for i := 0; i < n; i++ {
		s += " "
	}
	return s
}

func fullJustify(words []string, maxWidth int) []string {
	var result []string
	var line []string
	var lineWidth, wordWidth int
	for _, word := range words {
		if lineWidth > 0 && lineWidth+len(word)+1 > maxWidth {
			restWidth := maxWidth - wordWidth
			if len(line) == 1 {
				result = append(result, line[0]+getSpace(restWidth))
			} else {
				avgSpaceNum := restWidth / (len(line) - 1)
				extraSpaceNum := restWidth - avgSpaceNum*(len(line)-1)
				row := ""
				avgSpace := getSpace(avgSpaceNum)
				for i, w := range line {
					row += w
					if i == len(line)-1 {
						break
					}
					row += avgSpace
					if extraSpaceNum > 0 {
						row += " "
						extraSpaceNum--
					}
				}
				result = append(result, row)
			}
			line = []string{}
			lineWidth = 0
			wordWidth = 0
		}

		line = append(line, word)
		if lineWidth == 0 {
			lineWidth += len(word)
		} else {
			lineWidth += len(word) + 1
		}
		wordWidth += len(word)
	}
	if len(line) > 0 {
		restWidth := maxWidth - lineWidth
		row := strings.Join(line, " ") + getSpace(restWidth)
		result = append(result, row)
	}
	return result
}
