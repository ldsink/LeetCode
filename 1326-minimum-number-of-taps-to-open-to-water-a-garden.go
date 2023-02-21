package main

func minTaps(n int, ranges []int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	taps := make([]int, n+1)
	for i := 0; i <= n; i++ {
		if ranges[i] == 0 {
			continue
		}
		if left, right := i-ranges[i], min(n, i+ranges[i]); left <= 0 { // 当前水龙头可以覆盖起点
			taps[right] = 1
		} else {
			for ; left < right; left++ {
				if taps[left] == 0 {
					continue
				}
				if taps[right] == 0 || taps[right] > taps[left]+1 {
					taps[right] = taps[left] + 1
				}
			}
		}
	}
	if taps[n] == 0 {
		return -1
	}
	return taps[n]
}
