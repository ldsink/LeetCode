package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sc := make(map[rune]int)
	tc := make(map[rune]int)
	for _, r := range s {
		sc[r] ++
	}
	for _, r := range t {
		tc[r] ++
	}
	for r, c := range sc {
		if c != tc[r] {
			return false
		}
	}
	return true
}
