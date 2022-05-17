package main

func isLTE(a, b string, order map[rune]int) bool {
	j := len(a)
	if j > len(b) {
		j = len(b)
	}
	for i := 0; i < j; i++ {
		var idxA, idxB int
		if v, ok := order[rune(a[i])]; ok {
			idxA = v
		}
		if v, ok := order[rune(b[i])]; ok {
			idxB = v
		}
		if idxA < idxB {
			return true
		} else if idxA > idxB {
			return false
		}
	}
	return len(a) <= len(b)
}

func isAlienSorted(words []string, order string) bool {
	o := make(map[rune]int)
	for i, r := range order {
		o[r] = i + 1
	}
	for i := 0; i < len(words)-1; i++ {
		if isLTE(words[i], words[i+1], o) {
			continue
		}
		return false
	}
	return true
}
