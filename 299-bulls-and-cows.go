package main

import "fmt"

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func getHint(secret string, guess string) string {
	var bull, cow int
	var expect, actual [10]int
	for i := 0; i < len(guess); i++ {
		if secret[i] == guess[i] {
			bull++
			continue
		}
		expect[secret[i]-'0']++
		actual[guess[i]-'0']++
	}
	for i := 0; i < 10; i++ {
		cow += min(expect[i], actual[i])
	}
	return fmt.Sprintf("%dA%dB", bull, cow)
}
