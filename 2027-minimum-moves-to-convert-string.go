package main

func minimumMoves(s string) int {
	var result int
	for i := 0; i < len(s); i++ {
		if s[i] == 'X' {
			result++
			i += 2
		}
	}
	return result
}
