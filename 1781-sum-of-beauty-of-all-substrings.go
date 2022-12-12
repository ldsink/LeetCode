package main

import "math"

func beautySum(s string) int {
	const L = 26
	const A = 'a'
	getBeauty := func(start, end int) int {
		count := make([]int, L)
		for i := start; i < end; i++ {
			count[s[i]-A]++
		}
		max, min := 0, math.MaxInt
		for i := 0; i < L; i++ {
			if count[i] == 0 {
				continue
			}
			if max < count[i] {
				max = count[i]
			}
			if min > count[i] {
				min = count[i]
			}
		}
		return max - min
	}
	var result int
	for i := 0; i < len(s); i++ {
		for j := i + 2; j <= len(s); j++ {
			result += getBeauty(i, j)
		}
	}
	return result
}
