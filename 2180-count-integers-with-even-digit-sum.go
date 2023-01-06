package main

import "strconv"

func countEven(num int) int {
	var result int
	for i := 0; i <= num; i += 10 {
		s := strconv.Itoa(i)
		var sum int
		for _, c := range s {
			sum += int(c - '0')
		}
		even := 0
		if sum%2 == 0 {
			even = 1
		}
		odd := 1 ^ even
		for j := 0; j < 10 && i+j <= num; j++ {
			if j%2 == 0 {
				result += even
			} else {
				result += odd
			}
		}
	}
	return result - 1
}
