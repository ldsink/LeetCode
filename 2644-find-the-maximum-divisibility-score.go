package main

import "sort"

func maxDivScore(nums []int, divisors []int) int {
	sort.Ints(divisors)
	var count, current, result int
	count = -1
	for _, divisor := range divisors {
		current = 0
		for _, num := range nums {
			if num%divisor == 0 {
				current++
			}
		}
		if count < current {
			count = current
			result = divisor
		}
	}
	return result
}
