package main

func secondHighest(s string) int {
	first := -1
	second := -1
	for _, c := range s {
		if '0' <= c && c <= '9' {
			val := int(c - '0')
			if first < val {
				first, second = val, first
			} else if first != val && second < val {
				second = val
			}
		}
	}
	return second
}
