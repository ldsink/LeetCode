package main

func getMinCut(start, end int, s string, count int, palindrome, minCost [][]int) bool {
	if palindrome[start][end] == 0 {
		palindrome[start][end] = 1
		for i, j := start, end; i < j; i, j = i+1, j-1 {
			if s[i] != s[j] {
				palindrome[start][end] = 2
				break
			}
		}
	}

	if count == 0 {
		return palindrome[start][end] == 1
	}

	if minCost[start][end] != -1 && count <= minCost[start][end] {
		return false
	}
	minCost[start][end] = count

	for i := start; i < end; i++ {
		if getMinCut(start, i, s, 0, palindrome, minCost) && getMinCut(i+1, end, s, count-1, palindrome, minCost) {
			return true
		}
	}
	return false
}

func minCut(s string) int {
	if len(s) <= 1 {
		return 0
	}
	var palindrome, minCost [][]int
	for i := 0; i < len(s); i++ {
		palindrome = append(palindrome, make([]int, len(s)))
		minCost = append(minCost, make([]int, len(s)))
		for j := i + 1; j < len(s); j++ {
			minCost[i][j] = -1
		}
	}
	var result int
	for ; ; result++ {
		if getMinCut(0, len(s)-1, s, result, palindrome, minCost) {
			break
		}
	}
	return result
}
