package main

import "sort"

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	var result int
	for i, j := 0, len(people)-1; i <= j; {
		if people[i]+people[j] <= limit {
			i, j = i+1, j-1
		} else {
			j = j - 1
		}
		result++
	}
	return result
}
