package main

func longestDecomposition(text string) int {
	result := make(map[int]int)
	var f func(int, int) int
	f = func(start int, end int) int {
		if start >= end {
			return 0
		}
		key := start<<10 | end
		if k, ok := result[key]; ok {
			return k
		}
		k := 1
		for l := 1; l <= (end-start)>>1; l++ {
			if text[start:start+l] == text[end-l:end] {
				if v := 2 + f(start+l, end-l); k < v {
					k = v
				}
			}
		}
		result[key] = k
		return k
	}
	return f(0, len(text))
}
