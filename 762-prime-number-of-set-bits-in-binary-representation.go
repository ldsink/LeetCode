package main

import "math/bits"

func countPrimeSetBits(left int, right int) int {
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61}
	isPrime := make(map[int]bool)
	for _, num := range primes {
		isPrime[num] = true
	}
	var result int
	for i := uint(left); i <= uint(right); i++ {
		if isPrime[bits.OnesCount(i)] {
			result++
		}
	}
	return result
}
