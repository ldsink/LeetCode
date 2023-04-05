package main

func commonFactors(a int, b int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	var result int
	for i, j := 1, min(a, b); i <= j; i++ {
		if a%i == 0 && b%i == 0 {
			result++
		}
	}
	return result
}
