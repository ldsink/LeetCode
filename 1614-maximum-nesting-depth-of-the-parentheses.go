package main

func maxDepth(s string) int {
	var depth, max int
	for _, r := range []rune(s) {
		if r == '(' {
			depth++
			if max < depth {
				max++
			}
		} else if r == ')' {
			depth--
		}
	}
	return max
}
