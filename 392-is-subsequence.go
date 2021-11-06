package main

import "strings"

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	if len(s) > len(t) {
		return false
	}
	if len(s) == len(t) {
		return s == t
	}
	if len(s) == 1 {
		return strings.Contains(t, s)
	}
	for i := 0; i < len(t); i++ {
		if t[i] == s[0] {
			return isSubsequence(s[1:], t[i+1:])
		}
	}
	return false
}
