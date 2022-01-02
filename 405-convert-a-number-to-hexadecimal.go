package main

import (
	"strings"
)

func getHexStr(arr [32]bool) string {
	var rs []rune
	charSet := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}
	for i := 0; i < 32; i += 4 {
		idx := 0
		for j := 0; j < 4; j++ {
			if arr[i+j] {
				idx += 1 << j
			}
		}
		rs = append(rs, charSet[idx])
	}
	for i := 0; i < 4; i++ {
		rs[i], rs[7-i] = rs[7-i], rs[i]
	}
	return strings.TrimLeft(string(rs), "0")
}

func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	var arr [32]bool
	if num > 0 {
		for i := 0; i < 32; i++ {
			if num%2 == 1 {
				arr[i] = true
			}
			num >>= 1
		}
	} else {
		num = -num
		// 先计算原码
		for i := 0; i < 32; i++ {
			if num%2 == 1 {
				arr[i] = true
			}
			num >>= 1
		}
		// 原码取反之后加一
		for i := 0; i < 32; i++ {
			arr[i] = !arr[i]
		}
		carry := true
		for i := 0; carry && i < 32; i++ {
			if arr[i] {
				arr[i] = false
			} else {
				carry = false
				arr[i] = true
			}
		}
		arr[31] = true // 符号位表示负数
	}
	return getHexStr(arr)
}
