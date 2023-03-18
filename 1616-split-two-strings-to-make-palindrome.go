package main

func checkPalindromeFormation(a string, b string) bool {
	if len(a)&1 == 0 {
		a = a[:len(a)>>1] + "-" + a[len(a)>>1:]
		b = b[:len(b)>>1] + "-" + b[len(b)>>1:]
	}
	check := func(x, y string) bool {
		var middleLen int
		middle := len(x) >> 1
		for middleLen = 1; middle-middleLen >= 0 && x[middle-middleLen] == x[middle+middleLen]; middleLen++ {
		}
		middleLen--
		if middleLen == middle {
			return true
		}
		for i := 0; i < middle; i++ {
			if x[i] == y[len(y)-1-i] {
				if i+middleLen+1 == middle {
					return true
				}
			} else {
				break
			}
		}
		for i := 0; i < middle; i++ {
			if y[i] == x[len(y)-1-i] {
				if i+middleLen+1 == middle {
					return true
				}
			} else {
				break
			}
		}
		return false
	}
	return check(a, b) || check(b, a)
}
