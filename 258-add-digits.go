package main

func addDigits(num int) int {
	if num < 9 {
		return num
	}
	if n := num % 9; n == 0 {
		return 9
	} else {
		return n
	}
}
