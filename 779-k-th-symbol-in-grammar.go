package main

func kthGrammar(n int, k int) int {
	var grammar int
	for k--; n > 1; n-- {
		if k%2 != 0 {
			grammar ^= 1
		}
		k >>= 1
	}
	return grammar
}
