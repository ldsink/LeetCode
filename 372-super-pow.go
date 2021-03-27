package main

func superPow(a int, b []int) int {
	if a == 1 || a == 0 {
		return a
	}
	if len(b) == 1 && b[0] == 0 {
		return 1
	}

	var values []int
	for ; len(b) > 0; {
		over := 0
		for i := 0; i < len(b); i++ {
			b[i] += 10 * over
			over = b[i] % 2
			b[i] /= 2
		}
		values = append(values, over)
		for i := 0; i < len(b) && b[i] == 0; {
			b = b[1:]
		}
	}

	result := 1
	for i := len(values) - 1; i >= 0; i-- {
		if values[i] == 1 {
			result = result * result * a
		} else {
			result = result * result
		}
		result %= 1337
	}
	return result % 1337
}
