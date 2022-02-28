package main

import (
	"math/bits"
)

func maximumRequests(n int, requests [][]int) int {
	var result int
	for idx := 1<<len(requests) - 1; idx > 0; idx-- {
		in := make([]int, n)
		out := make([]int, n)
		for i, r := range requests {
			if 1<<i&idx > 0 {
				out[r[0]]++
				in[r[1]]++
			}
		}
		var error bool
		for i := 0; i < n; i++ {
			if out[i] != in[i] {
				error = true
				break
			}
		}
		if error {
			continue
		}
		if c := bits.OnesCount(uint(idx)); result < c {
			result = c
		}
	}
	return result
}
