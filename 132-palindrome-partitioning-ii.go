package main

func minCut(s string) int {
	if len(s) <= 1 {
		return 0
	}
	// init min cost
	minCost := make([]int, len(s)+1)
	for i := 0; i <= len(s); i++ {
		minCost[i] = i
	}
	for i := 0; i < len(s); i++ {
		// odd
		for j, k := i, i; 0 <= j && k < len(s); j, k = j-1, k+1 {
			if s[j] != s[k] {
				break
			}
			if minCost[k+1] > minCost[j]+1 {
				minCost[k+1] = minCost[j] + 1
			}
		}
		// even
		for j, k := i, i+1; 0 <= j && k < len(s); j, k = j-1, k+1 {
			if s[j] != s[k] {
				break
			}
			if minCost[k+1] > minCost[j]+1 {
				minCost[k+1] = minCost[j] + 1
			}
		}
	}
	return minCost[len(s)] - 1
}
