package main

func checkPowersOfThree(n int) bool {
	var i int
	for i = 1; i < n; i *= 3 {
	}
	for ; n > 0 && i > 0; i /= 3 {
		if n >= i {
			n -= i
		}
	}
	return n == 0
}
