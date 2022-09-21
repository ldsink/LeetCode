package main

import (
	"math"
)

func kSimilarity(s1 string, s2 string) int {
	if s1 == s2 {
		return 0
	}
	var rs1, rs2 []rune
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			rs1 = append(rs1, rune(s1[i]))
			rs2 = append(rs2, rune(s2[i]))
		}
	}
	s1, s2 = string(rs1), string(rs2)

	const width = 3
	const mask = 1<<width - 1

	getValue := func(c uint8) int {
		return int(c - 'a')
	}
	getValueByOffset := func(val, offset int) int {
		return (val >> offset) & mask
	}
	sToi := func(s string) (n int) { // 压缩字符串内容到一个 int 内
		for i := 0; i < len(s); i++ {
			n ^= getValue(s[i]) << (i * width)
		}
		return
	}
	setZero := func(val, offset int) int { // 从 offset 开始设置 width 个 bit 为 0
		return ((val >> (offset + width)) << (offset + width)) ^ (val & ((1 << offset) - 1))
	}
	swap := func(val, offset1, offset2 int) int { // offset1 always greater than offset2
		a, b := getValueByOffset(val, offset1), getValueByOffset(val, offset2)
		val = setZero(setZero(val, offset1), offset2)
		val ^= b<<offset1 ^ a<<offset2
		return val
	}

	cache, target := make(map[int]int), sToi(s2)
	var search func(int) int
	search = func(s int) int {
		if s == target {
			return 0
		}
		if c, ok := cache[s]; ok {
			return c
		}
		cache[s] = math.MaxInt
		for i := (len(s1) - 1) * width; i > 0; i -= width { // swap from left to right
			if s>>i == target>>i {
				continue
			}
			targetVal := getValueByOffset(target, i)
			for j := i - width; j >= 0; j -= width { // find valid swap
				if getValueByOffset(s, j) == targetVal {
					if r := search(swap(s, i, j)); cache[s] > r+1 {
						cache[s] = r + 1
					}
				}
			}
			break
		}
		return cache[s]
	}
	return search(sToi(s1))
}
