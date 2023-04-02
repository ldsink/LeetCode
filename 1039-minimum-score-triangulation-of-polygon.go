package main

import (
	"fmt"
)

func minScoreTriangulation(values []int) int {
	getVal := func(a, b, c int) int {
		return values[a] * values[b] * values[c]
	}
	var full func(int, int) int     // 两个点之间完全切割完的情况
	var sep func(int, int, int) int // 以某个点为顶点，剩下的区域分隔的情况
	fCache, sCache := make(map[int]int), make(map[int]int)
	full = func(start int, end int) int {
		key := (start << 10) ^ end
		if v, ok := fCache[key]; ok {
			return v
		}
		if start+2 == end { // 已经相连，只有一种情况
			v := getVal(start, start+1, end)
			fCache[key] = v
			return v
		}
		v := full(start+1, end) + getVal(start, start+1, end)
		if t := full(start, end-1) + getVal(start, end-1, end); v > t {
			v = t
		}
		for middle := start + 2; middle <= end-2; middle++ {
			if t := getVal(start, end, middle) + full(start, middle) + full(middle, end); v > t {
				v = t
			}
		}
		fCache[key] = v
		return v
	}
	sep = func(p int, start int, end int) int {
		if start == end {
			return 0
		}
		key := (p << 20) ^ (start << 10) ^ end
		if v, ok := sCache[key]; ok {
			return v
		}
		if start+1 == end {
			v := getVal(p, start, end)
			sCache[key] = v
			return v
		}
		v := full(start, end) + getVal(p, start, end)
		for middle := start + 1; middle < end; middle++ {
			if t := sep(p, start, middle) + sep(p, middle, end); v > t {
				v = t
			}
		}
		sCache[key] = v
		return v
	}
	return full(0, len(values)-1)
}

func main() {
	fmt.Println(88, minScoreTriangulation([]int{5, 3, 5, 5, 1, 6, 2, 3}))
}
