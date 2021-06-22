package main

import (
	"sort"
	"strconv"
)

func lexicalOrder(n int) []int {
	var s []string
	for i := 1; i <= n; i++ {
		s = append(s, strconv.Itoa(i))
	}
	sort.Strings(s)
	var result []int
	for _, a := range s {
		i, _ := strconv.Atoi(a)
		result = append(result, i)
	}
	return result
}
