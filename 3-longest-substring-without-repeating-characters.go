package main

import "fmt"

type runeIndex map[rune]int

func removeRuneByIndex(index int, runeIndexP *runeIndex) {
	m := *runeIndexP
	for r, i := range m {
		if i <= index {
			delete(m, r)
		}
	}
}

func lengthOfLongestSubstring(s string) int {
	rs := []rune(s)
	m := make(runeIndex)
	var maxLength, length int
	for i, r := range rs {
		index := i + 1
		prevIndex := m[r]
		if prevIndex == 0 {
			length++
		} else {
			removeRuneByIndex(prevIndex, &m)
			length = index - prevIndex
		}
		m[r] = index
		if length > maxLength {
			maxLength = length
		}
	}
	return maxLength
}

func main() {
	var result int
	var s string

	s = "abcabcbb"
	result = lengthOfLongestSubstring(s)
	fmt.Println(result)

	s = "bbbbb"
	result = lengthOfLongestSubstring(s)
	fmt.Println(result)

	s = "pwwkew"
	result = lengthOfLongestSubstring(s)
	fmt.Println(result)
}
