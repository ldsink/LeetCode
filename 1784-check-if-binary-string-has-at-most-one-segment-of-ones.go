package main

func checkOnesSegment(s string) bool {
	var i int
	for i = 0; i < len(s); i++ {
		if s[i] != '1' {
			break
		}
	}
	for j := i + 1; j < len(s); j++ {
		if s[j] == '1' {
			return false
		}
	}
	return true
}
