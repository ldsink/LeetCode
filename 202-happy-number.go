package main

import (
	"strconv"
)

func mapper(r int32) int {
	if r == '0' {
		return 0
	} else if r == '1' {
		return 1
	} else if r == '2' {
		return 4
	} else if r == '3' {
		return 9
	} else if r == '4' {
		return 16
	} else if r == '5' {
		return 25
	} else if r == '6' {
		return 36
	} else if r == '7' {
		return 49
	} else if r == '8' {
		return 64
	} else {
		return 81
	}
}

func process(n int) int {
	s := strconv.Itoa(n)
	var result int
	for _, r := range s {
		result += mapper(r)
	}
	return result
}

func isHappy(n int) bool {
	nums := make(map[int]bool)
	for !nums[n] {
		nums[n] = true
		n = process(n)
		if n == 1 {
			return true
		}
	}
	return false
}
