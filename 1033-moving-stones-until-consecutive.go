package main

import "math"

func numMovesStones(a int, b int, c int) []int {
	getKey := func(a, b, c int) int {
		return a<<20 ^ b<<10 ^ c
	}
	result := make(map[int][2]int)
	var getMoves func(int, int, int) [2]int
	getMoves = func(a, b, c int) [2]int {
		k := getKey(a, b, c)
		if r, ok := result[k]; ok {
			return r
		}
		if a+1 == b && b+1 == c {
			result[k] = [2]int{}
			return result[k]
		}
		min, max := math.MaxInt, 0
		for i := a + 1; i < b; i++ {
			t := getMoves(i, b, c)
			if min > t[0] {
				min = t[0]
			}
			if max < t[1] {
				max = t[1]
			}
		}
		for i := b + 1; i < c; i++ {
			t := getMoves(b, i, c)
			if min > t[0] {
				min = t[0]
			}
			if max < t[1] {
				max = t[1]
			}
		}
		for i := a + 1; i < b; i++ {
			t := getMoves(a, i, b)
			if min > t[0] {
				min = t[0]
			}
			if max < t[1] {
				max = t[1]
			}
		}
		for i := b + 1; i < c; i++ {
			t := getMoves(a, b, i)
			if min > t[0] {
				min = t[0]
			}
			if max < t[1] {
				max = t[1]
			}
		}
		result[k] = [2]int{min + 1, max + 1}
		return result[k]
	}
	if a > b {
		a, b = b, a
	}
	if a > c {
		a, c = c, a
	}
	if b > c {
		b, c = c, b
	}
	r := getMoves(a, b, c)
	return r[:]
}
