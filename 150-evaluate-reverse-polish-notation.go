package main

import "strconv"

func evalRPN(tokens []string) int {
	var stack []int
	for _, token := range tokens {
		if token == "+" {
			stack[len(stack)-2] += stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		} else if token == "-" {
			stack[len(stack)-2] -= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		} else if token == "*" {
			stack[len(stack)-2] *= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		} else if token == "/" {
			stack[len(stack)-2] /= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		} else {
			n, _ := strconv.Atoi(token)
			stack = append(stack, n)
		}
	}
	return stack[0]
}
