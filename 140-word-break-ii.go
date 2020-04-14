package main

import (
	"sort"
)

type ByLen []string

func (a ByLen) Len() int {
	return len(a)
}

func (a ByLen) Less(i, j int) bool {
	return len(a[i]) < len(a[j])
}

func (a ByLen) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func _wordBreak(s string, wordDict []string, found []bool, result [][]string) []string {
	if found[len(s)] {
		return result[len(s)]
	}
	for _, word := range wordDict {
		if len(s) < len(word) {
			break
		}
		if s[:len(word)] == word {
			if len(s) == len(word) {
				result[len(s)] = append(result[len(s)], " "+word)
			} else if r := _wordBreak(s[len(word):], wordDict, found, result); len(r) > 0 {
				for _, w := range r {
					result[len(s)] = append(result[len(s)], " "+word+w)
				}
			}
		}
	}
	found[len(s)] = true
	return result[len(s)]
}

func wordBreak(s string, wordDict []string) []string {
	sort.Sort(ByLen(wordDict))
	found := make([]bool, len(s)+1)
	var result [][]string
	for i := 0; i <= len(s); i++ {
		result = append(result, []string{})
	}
	r := _wordBreak(s, wordDict, found, result)
	if len(r) == 0 {
		return []string{}
	}
	var words []string
	for _, w := range r {
		words = append(words, w[1:])
	}
	return words
}
