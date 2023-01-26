package main

func getSmallestString(n int, k int) string {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	runes := make([]rune, n)
	for i := 0; i < n; i++ {
		runes[i] = 'a'
	}
	delta := k - n
	for i := n - 1; delta > 0 && i >= 0; i-- {
		v := min(25, delta)
		delta -= v
		runes[i] += rune(v)
	}
	return string(runes)
}
