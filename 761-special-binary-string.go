package main

import (
	"fmt"
	"sort"
	"strings"
)

func makeLargestSpecial(s string) string {
	if len(s) <= 2 {
		return s
	}
	var subs []string
	count, start := 0, 0
	for idx, r := range s {
		if r == '1' {
			count++
		} else {
			count--
			if count == 0 {
				subs = append(subs, fmt.Sprintf("1%s0", makeLargestSpecial(s[start+1:idx])))
				start = idx + 1
			}
		}
	}
	sort.Slice(subs, func(i, j int) bool {
		return subs[i] > subs[j]
	})
	return strings.Join(subs, "")
}
