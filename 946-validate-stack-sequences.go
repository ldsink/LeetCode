package main

func validateStackSequences(pushed []int, popped []int) bool {
	var stack []int
	for _, num := range pushed {
		stack = append(stack, num)
		for len(stack) > 0 && stack[len(stack)-1] == popped[0] {
			stack = stack[:len(stack)-1]
			popped = popped[1:]
		}
	}
	for ; len(popped) > 0; stack, popped = stack[:len(stack)-1], popped[1:] {
		if stack[len(stack)-1] != popped[0] {
			return false
		}
	}
	return true
}
