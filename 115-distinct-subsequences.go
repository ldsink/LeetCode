package main

func numDistinct(s string, t string) int {
	var saved [][]int
	for i := 0; i < 2; i++ {
		var line []int
		for j := 0; j < len(t)+1; j++ {
			line = append(line, 0)
		}
		saved = append(saved, line)
	}
	m, n := len(s), len(t)
	saved[m%2][n] = 1
	for i := m - 1; i >= 0; i-- {
		k, l := i%2, (i+1)%2
		saved[k][n] = 1
		for j := n - 1; j >= 0; j-- {
			saved[k][j] = saved[l][j]
			if s[i] == t[j] {
				saved[k][j] += saved[l][j+1]
			}
		}
	}
	return saved[0][0]
}
