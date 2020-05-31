package main

import (
	"strconv"
)

func getExpressions(num string) []string {
	var ret []string

	n, err := strconv.Atoi(num)
	if err == nil && strconv.Itoa(n) == num {
		ret = append(ret, num)
	}

	for i := 1; i < len(num); i++ {
		leftNum, err := strconv.Atoi(num[:i])
		if err != nil || strconv.Itoa(leftNum) != num[:i] {
			break
		}
		if r := getExpressions(num[i:]); len(r) > 0 {
			for _, e := range r {
				ret = append(ret, num[:i]+"+"+e)
				ret = append(ret, num[:i]+"-"+e)
				ret = append(ret, num[:i]+"*"+e)
			}
		}
	}
	return ret
}

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

func addOperators(num string, target int) []string {
	var ret []string
	expressions := getExpressions(num)
	for _, expression := range expressions {
		if target == calculate(expression) {
			ret = append(ret, expression)
		}
	}
	return ret
}
