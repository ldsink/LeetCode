package main

func shortestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	var length, i, j int
	for length = len(s); ; length++ {
		i = length>>1 - 1 - (length - len(s))
		if length%2 == 0 {
			j = i + 1
		} else {
			j = i + 2
		}
		palindrome := true
		for ; i >= 0; i, j = i-1, j+1 {
			if s[i] != s[j] {
				palindrome = false
				break
			}
		}
		if palindrome {
			break
		}
	}
	var prefix []rune
	for i, j := 1, length-len(s); i <= j; i++ {
		prefix = append(prefix, rune(s[len(s)-i]))
	}
	return string(prefix) + s
}
