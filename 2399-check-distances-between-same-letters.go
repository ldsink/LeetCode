package main

func checkDistances(s string, distance []int) bool {
	idx := make(map[rune]int)
	for i, r := range s {
		if j, ok := idx[r]; !ok {
			idx[r] = i
		} else if i-j-1 != distance[r-'a'] {
			return false
		}
	}
	return true
}
