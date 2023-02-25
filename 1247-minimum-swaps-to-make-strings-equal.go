package main

func minimumSwap(s1 string, s2 string) int {
	var x, y int
	for i := 0; i < len(s1); i++ {
		if s1[i] == s2[i] {
			continue
		} else if s1[i] == 'x' {
			x++
		} else {
			y++
		}
	}
	if (x+y)%2 != 0 {
		return -1
	}
	return (x >> 1) + (y >> 1) + x&1 + y&1
}
