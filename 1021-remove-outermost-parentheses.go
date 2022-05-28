package main

func removeOuterParentheses(s string) string {
	var rs []rune
	var result string
	var start int
	for i, r := range []rune(s) {
		if len(rs) == 0 {
			start = i + 1
		}
		if r == ')' {
			rs = rs[:len(rs)-1]
			if len(rs) == 0 {
				result += s[start:i]
			}
		} else {
			rs = append(rs, r)
		}
	}
	return result
}
