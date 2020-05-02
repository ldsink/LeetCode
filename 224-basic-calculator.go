package main

import (
	"strconv"
)

func calculate(s string) int {
	var operator []rune
	var nums []int

	for i := 0; i < len(s); i++ {
		if '0' <= s[i] && s[i] <= '9' {
			j := i + 1
			for ; j < len(s) && '0' <= s[j] && s[j] <= '9'; j++ {
			}
			n, _ := strconv.Atoi(s[i:j])
			nums = append(nums, n)
			i = j - 1
		} else if s[i] == '(' {
			operator = append(operator, '(')
		} else if s[i] == '+' || s[i] == '-' || s[i] == ')' {
			for ; len(operator) > 0 && operator[len(operator)-1] != '('; {
				if operator[len(operator)-1] == '+' {
					nums[len(nums)-2] += nums[len(nums)-1]
				} else {
					nums[len(nums)-2] -= nums[len(nums)-1]
				}
				nums = nums[:len(nums)-1]
				operator = operator[:len(operator)-1]
			}
			if s[i] == ')' {
				operator = operator[:len(operator)-1]
			} else {
				operator = append(operator, rune(s[i]))
			}
		}
	}

	if len(operator) > 0 {
		if operator[len(operator)-1] == '+' {
			nums[0] += nums[1]
		} else {
			nums[0] -= nums[1]
		}
	}
	return nums[0]
}
