package main

import (
	"strings"
)

func evaluate(s string, knowledge [][]string) string {
	mapping := make(map[string]string)
	for _, item := range knowledge {
		mapping[item[0]] = item[1]
	}
	parts := strings.FieldsFunc(s, func(r rune) bool {
		return r == '(' || r == ')'
	})
	var elements []string
	var idx int
	for _, part := range parts {
		if strings.HasPrefix(s[idx:], part) {
			idx += len(part)
			elements = append(elements, part)
		} else {
			idx += len(part) + 2
			if val, ok := mapping[part]; ok {
				elements = append(elements, val)
			} else {
				elements = append(elements, "?")
			}
		}
	}
	return strings.Join(elements, "")
}
