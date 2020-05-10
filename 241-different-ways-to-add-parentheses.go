package main

import "strconv"

func _compute(s string, c map[string][]int) []int {
	if r, ok := c[s]; ok {
		return r
	}

	var operators []int
	for i, r := range s {
		if r == '+' || r == '-' || r == '*' {
			operators = append(operators, i)
		}
	}
	if len(operators) == 0 {
		n, _ := strconv.Atoi(s)
		c[s] = []int{n}
	} else if len(operators) == 1 {
		a, _ := strconv.Atoi(s[:operators[0]])
		b, _ := strconv.Atoi(s[operators[0]+1:])
		if o := s[operators[0]]; o == '+' {
			c[s] = []int{a + b}
		} else if o == '-' {
			c[s] = []int{a - b}
		} else {
			c[s] = []int{a * b}
		}
	} else {
		var result []int
		for _, i := range operators {
			left := _compute(s[:i], c)
			right := _compute(s[i+1:], c)
			if o := s[i]; o == '+' {
				for _, a := range left {
					for _, b := range right {
						result = append(result, a+b)
					}
				}
			} else if o == '-' {
				for _, a := range left {
					for _, b := range right {
						result = append(result, a-b)
					}
				}
			} else {
				for _, a := range left {
					for _, b := range right {
						result = append(result, a*b)
					}
				}
			}
		}
		c[s] = result
	}
	return c[s]
}

func diffWaysToCompute(input string) []int {
	cache := make(map[string][]int)
	return _compute(input, cache)
}
