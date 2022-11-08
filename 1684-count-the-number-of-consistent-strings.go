package main

func countConsistentStrings(allowed string, words []string) int {
	chars := make(map[rune]bool)
	for _, r := range allowed {
		chars[r] = true
	}
	var result int
	for _, word := range words {
		consistent := true
		for _, r := range word {
			if !chars[r] {
				consistent = false
				break
			}
		}
		if consistent {
			result++
		}
	}
	return result
}
