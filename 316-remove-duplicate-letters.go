package main

import (
	"sort"
)

var position map[rune][]int
var valid map[rune]bool
var result []rune
var runeSet []rune

func getLetters(index int) {
	if index == len(result) {
		return
	}
	for _, r := range runeSet {
		if !valid[r] {
			continue
		}
		start := position[r][0]
		upper := true
		for _, nr := range runeSet {
			if !valid[nr] || nr == r {
				continue
			}
			if position[nr][len(position[nr])-1] < start {
				upper = false
				break
			}
		}
		if !upper {
			continue
		}
		// 当前字符满足在所有的字符前面，一定是符合要求的
		result[index] = r
		delete(valid, r)
		delete(position, r)
		// 更新 position 内其他元素的数据
		for _, nr := range runeSet {
			if !valid[nr] {
				continue
			}
			var i, j int
			for i, j = 0, len(position[nr]); i < j; i++ {
				if position[nr][i] > start {
					j = i
					break
				}
			}
			position[nr] = position[nr][j:]
		}
		getLetters(index + 1)
		return
	}
}

func removeDuplicateLetters(s string) string {
	position = make(map[rune][]int)
	valid = make(map[rune]bool)
	for i, r := range []rune(s) {
		position[r] = append(position[r], i)
		valid[r] = true
	}
	result = make([]rune, len(position))
	for r, _ := range valid {
		runeSet = append(runeSet, r)
	}
	sort.Slice(runeSet, func(i, j int) bool {
		return runeSet[i] < runeSet[j]
	})
	getLetters(0)
	return string(result)
}
