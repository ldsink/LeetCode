package main

import "fmt"

func checkPalindrome(rs *[]rune, offset, length int) bool {
	for i, j, k := 0, length>>1, offset+length-1; i < j; i++ {
		if (*rs)[offset+i] != (*rs)[k-i] {
			return false
		}
	}
	return true
}

func longestPalindrome(s string) string {
	rs := []rune(s)
	if len(rs) == 0 {
		return ""
	}
	for i := len(rs); i > 1; i-- {
		for j := 0; j <= len(rs)-i; j++ {
			if checkPalindrome(&rs, j, i) {
				return string(rs[j : j+i])
			}
		}
	}
	return s[:1]
}

func main() {
	var s string
	s = "abcdbbfcba"
	fmt.Println("result: ", longestPalindrome(s))
}
