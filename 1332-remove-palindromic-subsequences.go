package main

func removePalindromeSub(s string) int {
	for i, j := 0, len(s)>>1; i < j; i++ {
		if s[i] != s[len(s)-1-i] {
			return 2
		}
	}
	return 1
}
