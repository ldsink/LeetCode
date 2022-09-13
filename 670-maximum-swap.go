package main

import (
	"sort"
	"strconv"
)

func maximumSwap(num int) int {
	origin := []rune(strconv.Itoa(num))
	sorted := []rune(strconv.Itoa(num))
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i] > sorted[j]
	})
	for i := 0; i < len(origin); i++ {
		if origin[i] != sorted[i] {
			for j := len(origin) - 1; j > i; j-- {
				if origin[j] == sorted[i] {
					origin[i], origin[j] = origin[j], origin[i]
					break
				}
			}
			break
		}
	}
	num, _ = strconv.Atoi(string(origin))
	return num
}
