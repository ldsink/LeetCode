package main

import (
	"strconv"
)

func power10(n int) int {
	result := 1
	for i := 0; i < n; i++ {
		result *= 10
	}
	return result
}

func countDigitOne(n int) int {
	if n <= 0 {
		return 0
	}
	num := strconv.Itoa(n)
	var count int
	for i := 0; i < len(num); i++ {
		// 受前面部分影响，满排列情况
		prefix, _ := strconv.Atoi(num[:i+1])
		t := (prefix + 8) / 10
		p := power10(len(num) - i - 1)
		count += t * p
		// 受后面部分影响，超出多少个
		if num[i] == '1' {
			count++
			if len(num[i+1:]) > 0 {
				n, _ := strconv.Atoi(num[i+1:])
				count += n
			}
		}
	}
	return count
}
