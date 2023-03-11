package main

import (
	"strconv"
)

func findLongestSubarray(array []string) []string {
	numOver := make(map[int]int)
	numOver[0] = -1
	var numCount, alphaCount, maxLen, index int
	for i, s := range array {
		if _, err := strconv.Atoi(s); err == nil {
			numCount++
		} else {
			alphaCount++
		}
		if start, ok := numOver[numCount-alphaCount]; !ok {
			numOver[numCount-alphaCount] = i
		} else if maxLen < i-start {
			index, maxLen = start+1, i-start
		}
	}
	if maxLen == 0 {
		return []string{}
	}
	return array[index : index+maxLen]
}
