package main

import "math/bits"

func isPowerOfFour(num int) bool {
	if num <= 0 {
		return false
	}
	length := bits.Len(uint(num))
	if (length-1)%2 != 0 {
		return false
	}
	return num == (1 << (length - 1))
}
