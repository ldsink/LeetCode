package main

import (
	"fmt"
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)
	for _, str := range strs {
		s := strings.Split(str, "")
		sort.Strings(s)
		group := strings.Join(s, "")
		if _, ok := groups[group]; ok {
			groups[group] = append(groups[group], str)
		} else {
			groups[group] = []string{str}
		}
	}
	var result [][]string
	for _, v := range groups {
		result = append(result, v)
		fmt.Println(v)
	}
	return result
}
