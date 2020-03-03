package main

import (
	"sort"
)

func generateParenthesis(n int) []string {
	parenthesis := [][]string{
		[]string{},
		[]string{"()"},
	}
	for i := 2; i <= n; i++ {
		current := make(map[string]bool)
		// for single
		for _, c := range parenthesis[i-1] {
			var condition string
			condition = "(" + c + ")"
			current[condition] = true
			condition = "()" + c
			current[condition] = true
			condition = c + "()"
			current[condition] = true
		}
		// other cases
		for j, k := 2, i>>1; j <= k; j++ {
			l := i - j
			for _, cj := range parenthesis[j] {
				for _, cl := range parenthesis[l] {
					var condition string
					condition = cj + cl
					current[condition] = true
					condition = cl + cj
					current[condition] = true
				}
			}

		}
		// map to slice
		array := make([]string, 0, len(current))
		for c := range current {
			array = append(array, c)
		}
		sort.Strings(array)
		parenthesis = append(parenthesis, array)
	}
	return parenthesis[n]
}
