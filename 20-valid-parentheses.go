package main

func isValid(s string) bool {
	rs := []rune(s)
	match := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	var stack []rune
	for _, r := range rs {
		if val, ok := match[r]; ok {
			if len(stack) == 0 {
				return false
			}
			if stack[len(stack)-1] != val {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, r)
		}
	}
	return len(stack) == 0
}
