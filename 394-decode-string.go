package main

import (
	"strconv"
	"strings"
)

func decodeString(s string) string {
	result := ""
	leftWrap := "["
	rightWrap := "]"
	for i := 0; i < len(s); i++ {
		if strings.Contains("1234567890", s[i:i+1]) {
			leftIndex := i + 1
			for ; s[leftIndex:leftIndex+1] != leftWrap; leftIndex++ {
			}
			dupCount, _ := strconv.Atoi(s[i:leftIndex])
			wrapCount := 1
			rightIndex := leftIndex + 1
			for ; ; rightIndex++ {
				if s[rightIndex:rightIndex+1] == leftWrap {
					wrapCount++
				} else if s[rightIndex:rightIndex+1] == rightWrap {
					wrapCount--
				}
				if wrapCount == 0 {
					break
				}
			}
			content := decodeString(s[leftIndex+1 : rightIndex])
			for ; dupCount > 0; dupCount-- {
				result += content
			}
			i = rightIndex
		} else {
			result += s[i : i+1]
		}
	}
	return result
}
