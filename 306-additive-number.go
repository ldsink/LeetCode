package main

import (
	"strconv"
	"strings"
)

func isValid(a, b, e string) bool {
	i, _ := strconv.Atoi(a)
	j, _ := strconv.Atoi(b)
	c := strconv.Itoa(i + j)
	if len(c) > len(e) {
		return false
	} else if !strings.HasPrefix(e, c) {
		return false
	}
	if e == c {
		return true
	}
	return isValid(b, c, e[len(c):])
}

func isAdditiveNumber(num string) bool {
	if len(num) < 3 {
		return false
	}
	for i := 1; i <= len(num)/2; i++ {
		if num[0] == '0' && i != 1 {
			continue
		}
		for j := 1; j <= len(num)/2; j++ {
			if num[i] == '0' && j != 1 {
				continue
			}
			if isValid(num[:i], num[i:i+j], num[i+j:]) {
				return true
			}
		}
	}
	return false
}
