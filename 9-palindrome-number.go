package main

import (
	"strconv"
)

func isPalindrome(x int) bool {
	s := strconv.FormatInt(int64(x), 10)
	rs := []rune(s)
	for i, j, k := 0, len(rs)>>1, len(rs)-1; i < j; i++ {
		if rs[i] != rs[k-i] {
			return false
		}
	}
	return true
}
