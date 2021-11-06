package main

func lastRemaining(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 || n == 3 {
		return 2
	}
	remain := n/2 + (n % 2)
	reverseN := n - remain
	reverseResult := lastRemaining(reverseN)
	currIdx := reverseN + 1 - reverseResult
	return 2 * currIdx
}
