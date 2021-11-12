package main

import "strings"

func makeFancyString(s string) string {
	var count int
	var c string
	var result []string
	for i := 0; i < len(s); i++ {
		if count == 0 {
			count++
			c = s[i : i+1]
		} else if s[i:i+1] == c[:1] {
			if count < 2 {
				c += s[i : i+1]
			}
			count++
		} else {
			result = append(result, c)
			count = 1
			c = s[i : i+1]
		}
	}
	if c != "" {
		result = append(result, c)
	}
	return strings.Join(result, "")
}
