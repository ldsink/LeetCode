package main

func countPrimes(n int) int {
	notPrime := make(map[int]bool)
	var count int
	for i := 2; i < n; i++ {
		if notPrime[i] {
			continue
		}
		count++
		for j := i << 1; j < n; j += i {
			notPrime[j] = true
		}
	}
	return count
}
