package main

import "strconv"

func maximumValue(strs []string) int {
	var result int
	for _, s := range strs {
		val, err := strconv.Atoi(s)
		if err != nil {
			val = len(s)
		}
		if result < val {
			result = val
		}
	}
	return result
}
