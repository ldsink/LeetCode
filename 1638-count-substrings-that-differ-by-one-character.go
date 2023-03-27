package main

func countSubstrings(s string, t string) int {
	var result int
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(t); j++ {
			canCount := false
			for k := 0; i+k < len(s) && j+k < len(t); k++ {
				if s[i+k] != t[j+k] {
					if canCount {
						break
					} else {
						canCount = true
						result++
					}
				} else if canCount {
					result++
				}
			}
		}
	}
	return result
}
