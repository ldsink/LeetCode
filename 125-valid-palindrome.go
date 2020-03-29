package main

import "strings"

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	var b []byte
	for i := 0; i < len(s); i++ {
		if ('a' <= s[i] && s[i] <= 'z') || ('0' <= s[i] && s[i] <= '9') {
			b = append(b, s[i])
		}
	}
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		if b[i] != b[j] {
			return false
		}
	}
	return true
}
