package main

func oneEditAway(first string, second string) bool {
	if len(first) > len(second) {
		first, second = second, first
	}
	if len(second)-len(first) > 1 {
		return false
	}
	var diff int
	if len(first) == len(second) {
		for i := 0; i < len(first) && diff < 2; i++ {
			if first[i] != second[i] {
				diff++
			}
		}
	} else {
		for i := 0; i < len(first) && diff < 2; i++ {
			if first[i] != second[i+diff] {
				diff++
				i--
			}
		}
	}
	return diff < 2
}
