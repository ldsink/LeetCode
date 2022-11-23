package main

import (
	"strconv"
	"strings"
)

func countBalls(lowLimit int, highLimit int) int {
	boxes := make([]int, 50)
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	for i := lowLimit; i <= highLimit; i++ {
		var sum int
		for _, part := range strings.Split(strconv.Itoa(i), "") {
			val, _ := strconv.Atoi(part)
			sum += val
		}
		n := min(10-i%10, highLimit-i+1)
		for j := 0; j < n; j++ {
			boxes[sum+j]++
		}
		i += n - 1
	}
	var result int
	for _, num := range boxes {
		if result < num {
			result = num
		}
	}
	return result
}
