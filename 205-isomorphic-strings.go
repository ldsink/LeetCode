package main

func isIsomorphic(s string, t string) bool {
	mapper := make(map[rune]rune)
	used := make(map[rune]bool)
	for i, c := range s {
		tr := rune(t[i])
		if r, ok := mapper[c]; ok {
			if r != tr {
				return false
			}
		} else if used[tr] {
			return false
		} else {
			mapper[c] = tr
			used[tr] = true
		}
	}
	return true
}
