package main

func isScramble(s1 string, s2 string) bool {
	if len(s1) == 1 {
		return s1 == s2
	}

	// prune
	var b1, b2 [26]int
	for i := 0; i < len(s1); i++ {
		b1[s1[i]-'a']++
		b2[s2[i]-'a']++
	}
	for i := 0; i < 26; i++ {
		if b1[i] != b2[i] {
			return false
		}
	}

	for i := 1; i < len(s1); i++ {
		rest := len(s1) - i
		r := (isScramble(s1[:i], s2[:i]) && isScramble(s1[i:], s2[i:])) || (isScramble(s1[:i], s2[rest:]) && isScramble(s1[i:], s2[:rest]))
		if r {
			return r
		}
	}
	return false
}
