package main

func titleToNumber(s string) int {
	var result int
	for i, j := len(s)-1, 1; i >= 0; i, j = i-1, j*26 {
		result += j * int(s[i]-'A'+1)
	}
	return result
}
