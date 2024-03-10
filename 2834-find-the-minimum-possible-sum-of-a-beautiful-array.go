package main

func minimumPossibleSum(n int, target int) int {
	middle := target >> 1
	const mod = 1e9 + 7
	if n <= middle {
		return ((n * (n + 1)) >> 1) % mod
	}
	over := n - middle
	result := ((middle * (middle + 1)) >> 1) % mod
	result += (target * over) % mod
	result += ((over * (over - 1)) >> 1) % mod
	return result % mod
}
