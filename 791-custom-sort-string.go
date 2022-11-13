package main

import "sort"

func customSortString(order string, s string) string {
	index := make([]int, 26)
	for i, r := range order {
		index[r-'a'] = i
	}
	result := []rune(s)
	sort.Slice(result, func(i, j int) bool {
		return index[result[i]-'a'] < index[result[j]-'a']
	})
	return string(result)
}
