package main

func maxPower(s string) int {
	var i, j, result int
	for i = 0; i < len(s); i++ {
		for j = i + 1; j < len(s) && s[i] == s[j]; j++ {
		}
		if result < j-i {
			result = j - i
		}
		i = j - 1
	}
	return result
}
