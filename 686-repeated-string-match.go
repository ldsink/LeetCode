package main

import "strings"

func repeatedStringMatch(a string, b string) int {
	// 首先构造出一个最长的串
	l := len(b) / len(a)
	if len(b)%len(a) > 0 {
		l++
	}
	max := l * 2
	var arr []string
	for i := 0; i < max; i++ {
		arr = append(arr, a)
	}
	c := strings.Join(arr, "")
	// 然后直接枚举长度
	for i := len(a); i <= len(c); i = i + len(a) {
		if strings.Contains(c[:i], b) {
			return i / len(a)
		}
	}
	return -1
}
