package main

func numberOfMatches(n int) int {
	var matches int
	for n > 1 {
		matches += n / 2
		n = n/2 + (n % 2)
	}
	return matches
}
