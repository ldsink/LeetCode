package main

func majorityElement(nums []int) int {
	var stack []int
	for _, num := range nums {
		if len(stack) == 0 {
			stack = append(stack, num)
		} else if num != stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, num)
		}
	}
	return stack[0]
}
