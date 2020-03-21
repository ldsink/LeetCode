package main

func myPow(x float64, n int) float64 {
	if n < 0 {
		return 1 / myPow(x, -n)
	}
	var result float64 = 1
	for p := x; n > 0; n = n >> 1 {
		if n&0x1 == 1 {
			result *= p
		}
		p = p * p
	}
	return result
}
