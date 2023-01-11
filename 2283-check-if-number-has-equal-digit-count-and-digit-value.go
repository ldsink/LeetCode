package main

import "strconv"

func digitCount(num string) bool {
	counter := make(map[int]int)
	for _, r := range num {
		v, _ := strconv.Atoi(string(r))
		counter[v]++
	}
	for i := 0; i < len(num); i++ {
		c, _ := strconv.Atoi(num[i : i+1])
		if counter[i] != c {
			return false
		}
	}
	return true
}
