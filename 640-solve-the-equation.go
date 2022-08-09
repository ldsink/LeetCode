package main

import (
	"fmt"
	"strconv"
	"strings"
)

func solveEquation(equation string) string {
	type Value struct {
		x   int
		num int
	}
	var getValue func(string) Value
	getValue = func(s string) Value {
		if len(s) == 0 {
			return Value{}
		}

		idx, sign := 0, 1
		if s[idx] == '-' {
			sign = -1
			idx++
		} else if s[idx] == '+' {
			idx++
		}

		isX, hasNext := false, false
		var num int
		for i := idx; i < len(s); i++ {
			if s[i] == '+' || s[i] == '-' || s[i] == 'x' {
				if s[i] == 'x' {
					isX = true
					if idx == i {
						num = 1
					} else {
						num, _ = strconv.Atoi(s[idx:i])
					}
					idx = i + 1
				} else {
					num, _ = strconv.Atoi(s[idx:i])
					idx = i
				}
				hasNext = true
				break
			}
		}
		var next Value
		if hasNext {
			next = getValue(s[idx:])
		} else {
			num, _ = strconv.Atoi(s[idx:])
		}
		if isX {
			next.x += num * sign
		} else {
			next.num += num * sign
		}
		return next
	}

	expressions := strings.SplitN(equation, "=", 2)
	left, right := getValue(expressions[0]), getValue(expressions[1])
	v := Value{x: left.x - right.x, num: right.num - left.num}
	if v.x == 0 {
		if v.num == 0 {
			return "Infinite solutions"
		} else {
			return "No solution"
		}
	}
	return fmt.Sprintf("x=%d", v.num/v.x)
}
