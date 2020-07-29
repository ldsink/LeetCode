package main

func isPowerOfThree(n int) bool {
	for ; ; n /= 3 {
		if n == 1 {
			return true
		}
		if n < 3 || n%3 != 0 {
			return false
		}
	}
}
