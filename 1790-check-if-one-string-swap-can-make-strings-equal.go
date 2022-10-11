package main

func areAlmostEqual(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}
	var a, b []rune
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			a, b = append(a, rune(s1[i])), append(b, rune(s2[i]))
		}
	}
	if len(a) != 2 {
		return false
	}
	return a[0] == b[1] && a[1] == b[0]
}
