package main

import (
	"strconv"
	"strings"
)

func areNumbersAscending(s string) bool {
	prev := -1
	parts := strings.Split(s, " ")
	for _, p := range parts {
		num, err := strconv.Atoi(p)
		if err != nil {
			continue
		}
		if prev < num {
			prev = num
		} else {
			return false
		}
	}
	return true
}
