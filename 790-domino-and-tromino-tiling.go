package main

func numTilings(n int) int {
	// f(n) = f(n-1) + f(n-2) + 2 * f(n-3) + 2 * f(n-4) + ... + 2 * f(1) + 2 * f(0)
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}
	const mod = 1000000007
	result := make([]int, n+1)
	result[0], result[1], result[2] = 1, 1, 2
	for i := 3; i <= n; i++ {
		result[i] = result[i-1] + result[i-2]
		for j := i - 3; j >= 0; j-- {
			result[i] += 2 * result[j]
		}
		result[i] %= mod
	}
	return result[n]
}
