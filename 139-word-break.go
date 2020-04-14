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

func _wordBreak(s string, wordDict []string, failed []bool) bool {
	if failed[len(s)] {
		return false
	}
	for _, word := range wordDict {
		if len(s) < len(word) {
			break
		}
		if s[:len(word)] == word {
			if len(s) == len(word) || _wordBreak(s[len(word):], wordDict, failed) {
				return true
			}
		}
	}
	failed[len(s)] = true
	return false
}

func wordBreak(s string, wordDict []string) bool {
	sort.Sort(ByLen(wordDict))
	failed := make([]bool, len(s)+1)
	return _wordBreak(s, wordDict, failed)
}
