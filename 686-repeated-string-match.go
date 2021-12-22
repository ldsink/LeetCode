package main

import "strings"

func repeatedStringMatch(a string, b string) int {
	var count int
	var c string
	for count*len(a) < len(b) {
		c += a
		count++
	}
	if strings.Contains(c, b) {
		return count
	}
	for end := count * 2; count <= end; {
		c += a
		count++
		if strings.Contains(c, b) {
			return count
		}
	}
	return -1
}
