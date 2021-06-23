package main

func firstUniqChar(s string) int {
	count := make(map[rune]int)
	for _, r := range []rune(s) {
		count[r] += 1
	}
	for i := 0; i < len(s); i++ {
		if count[rune(s[i])] == 1 {
			return i
		}
	}
	return -1
}
