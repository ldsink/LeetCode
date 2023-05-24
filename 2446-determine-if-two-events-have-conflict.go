package main

import (
	"strconv"
)

func haveConflict(event1 []string, event2 []string) bool {
	getTime := func(s string) int {
		hour, _ := strconv.Atoi(s[:2])
		minute, _ := strconv.Atoi(s[3:])
		return hour*60 + minute
	}
	start1, end1 := getTime(event1[0]), getTime(event1[1])
	start2, end2 := getTime(event2[0]), getTime(event2[1])
	if start1 == start2 {
		return true
	} else if start1 > start2 {
		start1, end1, start2, end2 = start2, end2, start1, end1
	}
	return end1 >= start2
}
