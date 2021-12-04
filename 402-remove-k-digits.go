package main

import "strings"

func trimNum(num string) string {
	num = strings.TrimLeft(num, "0")
	if len(num) == 0 {
		num = "0"
	}
	return num
}

func removeKdigits(num string, k int) string {
	if k == 0 {
		return num
	}
	if len(num) <= k {
		return "0"
	}

	for i := 0; i < len(num)-1; i++ {
		if num[i] > num[i+1] {
			num = num[:i] + num[i+1:]
			return removeKdigits(trimNum(num), k-1)
		}
	}

	return num[:len(num)-k]
}
