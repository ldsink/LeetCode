package main

func reversePrefix(word string, ch byte) string {
	var i int
	for i = 0; i < len(word); i++ {
		if word[i] == ch {
			break
		}
	}
	if i != len(word) {
		rs := []rune(word)
		for j := 0; j <= (i >> 1); j++ {
			rs[j], rs[i-j] = rs[i-j], rs[j]
		}
		word = string(rs)
	}
	return word
}
