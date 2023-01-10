package main

func reinitializePermutation(n int) int {
	var gcd, lcm func(int, int) int
	gcd = func(a, b int) int {
		if a%b == 0 {
			return b
		}
		return gcd(b, a%b)
	}
	lcm = func(a, b int) int {
		return a * b / gcd(a, b)
	}
	getRound := func(idx int) int {
		for i, j := 0, idx; ; i++ {
			if i != 0 && j == idx {
				return i
			}
			if j%2 == 0 {
				j >>= 1
			} else {
				j = (n >> 1) + ((j - 1) >> 1)
			}
		}
	}
	result := 1
	for i := 0; i < n; i++ {
		result = lcm(result, getRound(i))
	}
	return result
}
