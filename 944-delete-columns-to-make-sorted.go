package main

func minDeletionSize(strs []string) int {
	var result int
	for j := 0; j < len(strs[0]); j++ {
		for i := 0; i < len(strs)-1; i++ {
			if strs[i][j] > strs[i+1][j] {
				result++
				break
			}
		}
	}
	return result
}
