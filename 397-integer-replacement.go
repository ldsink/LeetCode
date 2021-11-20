package main

import (
	"math/bits"
)

func integerReplacement(n int) int {
	var result int
	for ; n != 1; result++ {
		if n&1 == 0 {
			n >>= 1 // num / 2
		} else if n == 3 {
			n--
		} else {
			num := 1
			for n&num == num {
				num <<= 1
				num ^= 1
			}
			if bits.Len(uint(num)) < 3 {
				n ^= 1 // num - 1
			} else {
				n ^= num // num + 1
			}
		}
	}
	return result
}
