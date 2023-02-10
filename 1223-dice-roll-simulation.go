package main

func dieSimulator(n int, rollMax []int) int {
	const mod = 1e9 + 7
	curr := [6][15]int{}
	total := [6]int{} // 本次出现 roll 点的总次数
	for i := 0; i < 6; i++ {
		curr[i][0] = 1
		total[i] = 1
	}
	for ; n > 1; n-- {
		prev := curr
		curr = [6][15]int{}
		for i := 0; i < 6; i++ {
			for j := 0; j < 6; j++ {
				if i != j {
					curr[i][0] = (curr[i][0] + total[j]) % mod
				}
			}
			for j := 1; j < rollMax[i]; j++ {
				curr[i][j] = prev[i][j-1]
			}
		}
		for i := 0; i < 6; i++ {
			total[i] = 0
			for j := 0; j < rollMax[i]; j++ {
				total[i] += curr[i][j]
			}
			total[i] %= mod
		}
	}
	var result int
	for i := 0; i < 6; i++ {
		result += total[i]
	}
	return result % mod
}
