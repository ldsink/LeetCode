package main

import (
	"math/bits"
)

func hasAlternatingBits(n int) bool {
	var num int
	for i := bits.Len64(uint64(n)) - 1; i >= 0; i -= 2 {
		num ^= (1 << i)
	}
	return num == n
}
