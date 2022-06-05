package main

func consecutiveNumbersSum(n int) int {
	var result int
	isMatch := func(l int) bool {
		if l%2 == 1 {
			return n%l == 0
		}
		return n%l != 0 && 2*n%l == 0
	}
	for l := 1; (l*(l+1))>>1 <= n; l++ {
		if isMatch(l) {
			result++
		}
	}
	return result
}
