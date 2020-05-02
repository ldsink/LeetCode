package main

import (
	"strconv"
)

func calculate(s string) int {
	var operators []rune
	var nums []int
	operatorOrder := map[rune]int{
		')': 0,
		'+': 10,
		'-': 10,
		'*': 20,
		'/': 20,
		'(': 30,
	}
	for i := 0; i < len(s); i++ {
		if '0' <= s[i] && s[i] <= '9' {
			j := i + 1
			for ; j < len(s) && '0' <= s[j] && s[j] <= '9'; j++ {
			}
			n, _ := strconv.Atoi(s[i:j])
			nums = append(nums, n)
			i = j - 1
		} else if s[i] == '(' {
			operators = append(operators, '(')
		} else if order, ok := operatorOrder[rune(s[i])]; ok {
			for ; len(operators) > 0 && operators[len(operators)-1] != '(' && operatorOrder[operators[len(operators)-1]] >= order; {
				if operators[len(operators)-1] == '+' {
					nums[len(nums)-2] += nums[len(nums)-1]
				} else if operators[len(operators)-1] == '-' {
					nums[len(nums)-2] -= nums[len(nums)-1]
				} else if operators[len(operators)-1] == '*' {
					nums[len(nums)-2] *= nums[len(nums)-1]
				} else if operators[len(operators)-1] == '/' {
					nums[len(nums)-2] /= nums[len(nums)-1]
				}
				nums = nums[:len(nums)-1]
				operators = operators[:len(operators)-1]
			}
			if s[i] == ')' {
				operators = operators[:len(operators)-1]
			} else {
				operators = append(operators, rune(s[i]))
			}
		}
	}

	for ; len(operators) > 0; operators = operators[:len(operators)-1] {
		i, j := len(nums)-2, len(nums)-1
		if operators[len(operators)-1] == '+' {
			nums[i] += nums[j]
		} else if operators[len(operators)-1] == '-' {
			nums[i] -= nums[j]
		} else if operators[len(operators)-1] == '*' {
			nums[i] *= nums[j]
		} else if operators[len(operators)-1] == '/' {
			nums[i] /= nums[j]
		}
		nums = nums[:j]
	}
	return nums[0]
}
