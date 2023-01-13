package main

func rearrangeCharacters(s string, target string) int {
	count := make(map[rune]int)
	for _, r := range target {
		count[r]++
	}
	source := make(map[rune]int)
	for _, r := range s {
		source[r]++
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	result := -1
	for k, v := range count {
		t := source[k] / v
		if t == 0 {
			return 0
		}
		if result == -1 {
			result = t
		} else {
			result = min(result, t)
		}
	}
	return result
}
