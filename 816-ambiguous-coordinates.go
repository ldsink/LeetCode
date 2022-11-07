package main

import (
	"fmt"
)

func ambiguousCoordinates(s string) []string {
	s = s[1 : len(s)-1]

	getNumbers := func(n string) []string {
		if len(n) == 1 {
			return []string{n}
		}
		var numbers []string
		if n[0] != '0' {
			numbers = append(numbers, n)
		}
		if n[len(n)-1] != '0' {
			for i := 1; i < len(n); i++ {
				if i > 1 && n[0] == '0' {
					break
				}
				numbers = append(numbers, fmt.Sprintf("%s.%s", n[:i], n[i:]))
			}
		}
		return numbers
	}

	var result []string
	for i := 1; i < len(s); i++ {
		a := getNumbers(s[:i])
		if len(a) == 0 {
			continue
		}
		b := getNumbers(s[i:])
		if len(b) == 0 {
			continue
		}
		for j := 0; j < len(a); j++ {
			for k := 0; k < len(b); k++ {
				result = append(result, fmt.Sprintf("(%s, %s)", a[j], b[k]))
			}
		}
	}
	return result
}
