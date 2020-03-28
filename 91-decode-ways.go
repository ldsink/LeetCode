package main

func numDecodings(s string) int {
	ways := make([]int, len(s)+1)
	ways[len(s)] = 1
	for i := len(s) - 1; i >= 0; i-- {
		if i+1 <= len(s) && '1' <= s[i] && s[i] <= '9' {
			ways[i] += ways[i+1]
		}
		if i+2 <= len(s) {
			if s[i] == '1' && '0' <= s[i+1] && s[i+1] <= '9' {
				ways[i] += ways[i+2]
			} else if s[i] == '2' && '0' <= s[i+1] && s[i+1] <= '6' {
				ways[i] += ways[i+2]
			}
		}
	}
	return ways[0]
}
