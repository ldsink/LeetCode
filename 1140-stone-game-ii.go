package main

func stoneGameII(piles []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	suffixSum := make([]int, len(piles)+1)
	for i := len(piles) - 1; i >= 0; i-- {
		suffixSum[i] = suffixSum[i+1] + piles[i]
	}
	cached := make([]map[int]int, len(piles))
	for i := 0; i < len(piles); i++ {
		cached[i] = make(map[int]int)
	}
	var dfs func(int, int) int
	dfs = func(idx, m int) int {
		if idx == len(piles) {
			return 0
		}
		if val, ok := cached[idx][m]; ok {
			return val
		}
		var val int
		for x := m << 1; x > 0; x-- {
			if idx+x > len(piles) {
				continue
			}
			rival := dfs(idx+x, max(m, x))
			stone := suffixSum[idx] - rival
			if val < stone {
				val = stone
			}
		}
		cached[idx][m] = val
		return val
	}
	return dfs(0, 1)
}
