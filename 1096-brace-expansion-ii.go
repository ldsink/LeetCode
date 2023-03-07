package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

func braceExpansionII(expression string) []string {
	combinePrefixAndSuffix := func(prefixes, suffixes []string) []string {
		var result []string
		for _, prefix := range prefixes {
			for _, suffix := range suffixes {
				result = append(result, fmt.Sprintf("%s%s", prefix, suffix))
			}
		}
		sort.Strings(result)
		return result
	}
	removeDup := func(x []string) []string {
		var result []string
		exists := make(map[string]bool)
		for _, s := range x {
			if exists[s] {
				continue
			}
			exists[s] = true
			result = append(result, s)
		}
		sort.Strings(result)
		return result
	}
	if len(expression) == 0 {
		return []string{""}
	}
	var idx int
	// 以字母开头，处理全部结果
	if unicode.IsLetter(rune(expression[0])) {
		for ; idx < len(expression); idx++ {
			if !unicode.IsLetter(rune(expression[idx])) {
				break
			}
		}
		return combinePrefixAndSuffix([]string{expression[:idx]}, braceExpansionII(expression[idx:]))
	}
	// 以括号开头，处理整个括号
	var leftCount, rightCount int
	leftCount++
	for idx++; idx < len(expression); idx++ {
		if expression[idx] == '{' {
			leftCount++
		} else if expression[idx] == '}' {
			rightCount++
			if rightCount == leftCount {
				idx++
				break
			}
		}
	}
	suffixes := braceExpansionII(expression[idx:])
	var prefixes []string
	if leftCount == 1 { // 中间没有其他的花括号，直接快速处理
		prefixes = removeDup(strings.Split(expression[1:idx-1], ","))
	} else {
		start := 1
		for i, count := 1, 0; i < idx-1; i++ {
			if expression[i] == '{' {
				count++
			} else if expression[i] == '}' {
				count--
			} else if expression[i] == ',' && count == 0 {
				prefixes = append(prefixes, braceExpansionII(expression[start:i])...)
				start = i + 1
			}
		}
		prefixes = append(prefixes, braceExpansionII(expression[start:idx-1])...)
		prefixes = removeDup(prefixes)
	}
	return combinePrefixAndSuffix(prefixes, suffixes)
}
