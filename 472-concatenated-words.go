package main

import (
	"sort"
	"strings"
)

func isConcatenatedWord(s string, prefixes [][]string, processed map[string]bool) bool {
	if r, ok := processed[s]; ok {
		return r
	}
	var result bool
	for _, prefix := range prefixes[s[0]-'a'] {
		if strings.HasPrefix(s, prefix) && isConcatenatedWord(s[len(prefix):], prefixes, processed) {
			result = true
			break
		}
	}
	processed[s] = result
	return result
}

func findAllConcatenatedWordsInADict(words []string) []string {
	// 按照长度从小到长排序
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
	prefixes := make([][]string, 26)
	processed := make(map[string]bool)
	processed[""] = true
	var result []string
	for _, word := range words {
		if len(word) == 0 {
			continue
		}
		if isConcatenatedWord(word, prefixes, processed) {
			result = append(result, word)
		}
		prefixes[word[0]-'a'] = append(prefixes[word[0]-'a'], word)
		processed[word] = true
	}
	return result
}
