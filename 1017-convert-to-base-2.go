package main

import (
	"strings"
)

func baseNeg2(n int) string {
	if n == 0 {
		return "0"
	}
	var binary [64]int
	for i := 0; n > 0; n, i = n/2, i+1 {
		binary[i] = n % 2
	}
	var result []rune
	for i := 0; i < len(binary); i++ {
		if binary[i] == 0 {
			result = append(result, '0')
		} else if binary[i] == 2 { // 进位
			result = append(result, '0')
			binary[i+1]++
		} else {
			result = append(result, '1')
			if i%2 == 1 { // 负数位置，实际值应该是后一位正数加上当前位负数，正数位置需要 +1
				binary[i+1]++
			}
		}
	}
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return strings.TrimLeft(string(result), "0")
}
