package main

func trailingZeroes(n int) int {
	var result int
	for i := 5; n/i != 0; i *= 5 {
		result += n / i
	}
	return result
}
