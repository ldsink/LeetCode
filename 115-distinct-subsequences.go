package main

var (
	T string
	S string
)

func dynamicSearch(si, ti int, saved [][]int) {
	if saved[si][ti] != -1 {
		return
	}
	saved[si][ti] = 0
	if len(T) == ti {
		saved[si][ti] = 1
		return
	} else if len(S[si:]) < len(T[ti:]) {
		return
	} else if len(S[si:]) == len(T[ti:]) {
		if S[si:] == T[ti:] {
			saved[si][ti] = 1
		}
		return
	}
	dynamicSearch(si+1, ti, saved)
	saved[si][ti] += saved[si+1][ti]
	if S[si] == T[ti] {
		dynamicSearch(si+1, ti+1, saved)
		saved[si][ti] += saved[si+1][ti+1]
	}
}

func numDistinct(s string, t string) int {
	S, T = s, t
	m, n := len(s)+1, len(t)+1
	var saved [][]int
	for i := 0; i < m; i++ {
		var line []int
		for j := 0; j < n; j++ {
			line = append(line, -1)
		}
		saved = append(saved, line)
	}
	dynamicSearch(0, 0, saved)
	return saved[0][0]
}
