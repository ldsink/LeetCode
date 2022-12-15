package main

import "strconv"

func getLucky(s string, k int) int {
	ctoi := func(c rune) (r int) {
		a := strconv.Itoa(int(c - 'a' + 1))
		for _, c := range a {
			r += int(c - '0')
		}
		return
	}
	itoi := func(i int) (r int) {
		for ; i != 0; i /= 10 {
			r += i % 10
		}
		return
	}
	var result int
	for i := 0; i < k; i++ {
		if i == 0 {
			for _, c := range s {
				result += ctoi(c)
			}
		} else {
			result = itoi(result)
		}
	}
	return result
}
