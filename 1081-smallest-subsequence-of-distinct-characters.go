package main

import "sort"

var position map[rune][]int
var result []rune
var runeSet []rune

func getLetters(index, start int) {
	if index == len(result) {
		return
	}
	for idx := 0; idx < len(runeSet); idx++ {
		r := runeSet[idx]

		// 二分找到第一个符合要求的坐标
		left, right := 0, len(position[r])-1
		for ; left != right; {
			middle := (left + right) >> 1
			if position[r][middle] < start {
				left = middle + 1
			} else {
				right = middle
			}
		}
		start := position[r][left]

		// 当前字符是否满足在所有的字符前面。如果满足，当前位置就是这个字符。
		upper := true
		for _, nr := range runeSet {
			if nr == r {
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
		result[index] = r
		runeSet = append(runeSet[:idx], runeSet[idx+1:]...)
		getLetters(index+1, start)
		return
	}
}

func smallestSubsequence(text string) string {
	position = make(map[rune][]int)
	for i, r := range []rune(text) {
		position[r] = append(position[r], i)
	}
	result = make([]rune, len(position))
	for r, _ := range position {
		runeSet = append(runeSet, r)
	}
	sort.Slice(runeSet, func(i, j int) bool {
		return runeSet[i] < runeSet[j]
	})
	getLetters(0, 0)
	return string(result)
}
