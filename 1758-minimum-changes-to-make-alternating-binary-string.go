package main

func minOperations(s string) int {
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	chars := [2]uint8{'0', '1'}
	getOperations := func(idx int) (count int) {
		for i := 0; i < len(s); i += 2 {
			if s[i] != chars[idx] {
				count++
			}
		}
		idx = (idx ^ 1) & 1
		for i := 1; i < len(s); i += 2 {
			if s[i] != chars[idx] {
				count++
			}
		}
		return
	}
	return min(getOperations(0), getOperations(1))
}
