package main

func superPow(a int, b []int) int {
	if a == 1 || a == 0 {
		return a
	}
	result := 1
	if b[len(b)-1]%2 == 1 {
		b[len(b)-1] -= 1
		result *= a
	}
	over := 0
	for i := 0; i < len(b); i++ {
		b[i] += 10 * over
		over = b[i] % 2
		b[i] /= 2
	}
	for i := 0; i < len(b) && b[i] == 0; {
		b = b[1:]
	}
	if len(b) > 0 {
		sub := superPow(a, b)
		result *= sub * sub
	}
	return result % 1337
}
