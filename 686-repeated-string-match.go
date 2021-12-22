package main

import "strings"

func repeatedStringMatch(a string, b string) int {
	// 先构造一个 c 字符串，填充长度到不小于 b
	l := 0
	var ca []string
	for ; l < len(b); l = l + len(a) {
		ca = append(ca, a)
	}
	c := strings.Join(ca, "")
	// 然后在 c 中寻找某个偏移量，可以完全匹配 b
	for offset := 0; offset < len(c); offset++ {
		matched := true
		for i := 0; i < len(b); i++ {
			if c[(i+offset)%len(c)] != b[i] {
				matched = false
				break
			}
		}
		// 如果找到了满足条件的，再计算次数
		if matched {
			length := offset % len(a)
			length += len(b)
			count := length / len(a)
			if length%len(a) > 0 {
				count++
			}
			return count
		}
	}
	return -1
}
