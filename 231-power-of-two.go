package main

import "math/bits"

func isPowerOfTwo(n int) bool {
	length := bits.Len(uint(n))
	if length == 0 {
		return false
	}
	return (1 << uint(length-1)) == n
}
