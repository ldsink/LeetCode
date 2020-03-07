package main

func longestValidParentheses(s string) int {
	longestLength := make(map[int]int)
	var result int
	var indexStack []int
	var runeStack []rune

	for i, r := range []rune(s) {
		if len(indexStack) == 0 || r == '(' {
			indexStack = append(indexStack, i)
			runeStack = append(runeStack, r)
			continue
		}
		// matched, pop rune
		if runeStack[len(runeStack)-1] == '(' {
			index := indexStack[len(indexStack)-1]
			length := (i - index) + 1
			length += longestLength[index-1]
			longestLength[i] = length
			if result < length {
				result = length
			}
			indexStack = indexStack[:len(indexStack)-1]
			runeStack = runeStack[:len(runeStack)-1]
		}
	}
	return result
}
