package main

import (
	"math/bits"
)

func rangeBitwiseAnd(m int, n int) int {
	var result int
	for i := bits.Len(uint(n)) - 1; i >= 0; i-- {
		if j, k := (m>>uint(i))&0x1, (n>>uint(i))&0x1; j == k {
			result |= j << uint(i)
			continue
		}
		break
	}
	return result
}
