package main

func getCount(s string) [26]int {
	var count [26]int
	for _, r := range []rune(s) {
		if 'a' <= r && r <= 'z' {
			count[r-'a']++
		} else if 'A' <= r && r <= 'Z' {
			count[r-'A']++
		}
	}
	return count
}

func isMatch(a, b [26]int) bool {
	for i := 0; i < 26; i++ {
		if a[i] > b[i] {
			return false
		}
	}
	return true
}

func shortestCompletingWord(licensePlate string, words []string) string {
	count := getCount(licensePlate)
	index := -1
	for idx, word := range words {
		c := getCount(word)
		if isMatch(count, c) {
			if index == -1 || len(word) < len(words[index]) {
				index = idx
			}
		}
	}
	return words[index]
}
