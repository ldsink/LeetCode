package main

func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	}
	count := 9
	for i, j := n, 9; i > 1; i, j = i-1, j-1 {
		count *= j
	}
	return count + countNumbersWithUniqueDigits(n-1)
}
