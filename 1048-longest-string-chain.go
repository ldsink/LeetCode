package main

import "sort"

func longestStrChain(words []string) int {
	lenMap := make(map[int][]string)
	longest := make(map[string]int)
	for _, word := range words {
		lenMap[len(word)] = append(lenMap[len(word)], word)
	}
	var lengths []int
	for l := range lenMap {
		lengths = append(lengths, l)
	}
	sort.Ints(lengths)
	isPredecessor := func(a, b string) bool {
		if len(a) > len(b) {
			a, b = b, a
		}
		if len(a)+1 != len(b) {
			return false
		}
		diff := 1
		for i := 0; i < len(a) && diff >= 0; i++ {
			if a[i] != b[i+(1-diff)] {
				diff--
				i--
			}
		}
		return diff >= 0
	}
	result := 1
	for i, l := range lengths {
		hasNext := i+1 < len(lengths) && l+1 == lengths[i+1] // 有下一层
		for _, word := range lenMap[l] {
			if longest[word] == 0 {
				longest[word] = 1
			}
			if result < longest[word] {
				result = longest[word]
			}
			if hasNext {
				for _, next := range lenMap[l+1] {
					if isPredecessor(word, next) && longest[next] < longest[word]+1 {
						longest[next] = longest[word] + 1
					}
				}
			}
		}
	}
	return result
}
