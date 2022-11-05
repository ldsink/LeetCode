package main

func parseBoolExpr(expression string) bool {
	expr := []rune(expression)

	and := func(parts []bool) bool {
		for _, part := range parts {
			if !part {
				return false
			}
		}
		return true
	}
	or := func(parts []bool) bool {
		for _, part := range parts {
			if part {
				return true
			}
		}
		return false
	}

	var parseExpr func(int) (bool, int)
	parseParts := func(idx int) ([]bool, int) {
		var length int
		var parts []bool
		for true {
			r, l := parseExpr(idx + length)
			parts = append(parts, r)
			length += l
			if expr[idx+length] == ')' {
				break
			} else {
				length++ // skip ','
			}
		}
		return parts, length
	}
	parseExpr = func(idx int) (bool, int) {
		if expr[idx] == 't' {
			return true, 1
		} else if expr[idx] == 'f' {
			return false, 1
		} else if expr[idx] == '!' {
			content, length := parseExpr(idx + 2)
			return !content, length + 3
		} else if expr[idx] == '&' {
			parts, length := parseParts(idx + 2) // &(
			return and(parts), length + 3
		}
		parts, length := parseParts(idx + 2) // !(
		return or(parts), length + 3
	}

	result, _ := parseExpr(0)
	return result
}
