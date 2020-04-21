package main

func scoreOfParentheses(S string) int {
	var stack []rune
	var value []int
	for _, s := range S {
		if s == '(' {
			stack = append(stack, s)
			continue
		}

		for i, j := len(stack)-1, 0; i >= 0; i-- {
			if stack[i] == '(' {
				stack[i] = 'A'
				if j == 0 {
					value = append(value, 1)
				} else {
					value = append(value, 2*j)
				}
				break
			}
			j += value[len(value)-1]
			value = value[:len(value)-1]
			stack = stack[:len(stack)-1]
		}
	}

	var result int
	for _, v := range value {
		result += v
	}
	return result
}
