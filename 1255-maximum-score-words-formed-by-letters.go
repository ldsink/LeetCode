package main

import (
	"math/bits"
)

func maxScoreWords(words []string, letters []byte, score []int) int {
	const A = 'a'
	// 先统计字母表的情况
	lettersCount := make([]uint, 26)
	for _, r := range letters {
		lettersCount[r-A]++
	}
	type Letter struct {
		offset, mask int
	}
	letterMap := make([]Letter, 26)
	for i, offset := 0, 0; i < 26; i++ {
		l := bits.Len(lettersCount[i]) + 1
		letterMap[i] = Letter{offset: offset, mask: (1 << l) - 1}
		offset += l
	}
	// 再处理单词的情况
	type Word struct {
		score, value int
	}
	var wordList []Word
	for _, word := range words {
		count := make([]uint, 26)
		for _, r := range word {
			count[r-A]++
		}
		var s, v int
		valid := true
		for i := 0; i < 26; i++ {
			if count[i] > lettersCount[i] {
				valid = false
				break
			}
			v += (1 << letterMap[i].offset) * int(count[i])
			s += int(count[i]) * score[i]
		}
		if valid {
			wordList = append(wordList, Word{score: s, value: v})
		}
	}
	if len(wordList) == 0 { // 一个单词都凑不成，直接返回 0
		return 0
	}
	// 开始动态规划搜索
	canUseWord := func(l, w int) bool {
		for _, c := range letterMap {
			if ((l >> c.offset) & c.mask) < ((w >> c.offset) & c.mask) {
				return false
			}
		}
		return true
	}
	cached := make([]map[int]int, len(wordList))
	for i := 0; i < len(wordList); i++ {
		cached[i] = make(map[int]int)
	}
	var dfs func(int, int) int
	dfs = func(idx, val int) int {
		if idx == len(wordList) {
			return 0
		}
		if s, ok := cached[idx][val]; ok {
			return s
		}
		s := dfs(idx+1, val)
		if canUseWord(val, wordList[idx].value) {
			if t := wordList[idx].score + dfs(idx+1, val-wordList[idx].value); s < t {
				s = t
			}
		}
		cached[idx][val] = s
		return s
	}
	var letterValue int
	for i := 0; i < 26; i++ {
		letterValue += (1 << letterMap[i].offset) * int(lettersCount[i])
	}
	return dfs(0, letterValue)
}
