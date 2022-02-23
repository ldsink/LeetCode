package main

import "fmt"

func reverseOnlyLetters(s string) string {
	rs := []rune(s)
	var otherLetterIndexes []int
	var otherLetterRunes []rune
	var englishLetters []rune
	for i, r := range rs {
		if ('a' <= r && r <= 'z') || ('A' <= r && r <= 'Z') {
			englishLetters = append(englishLetters, r)
		} else {
			otherLetterIndexes = append(otherLetterIndexes, i)
			otherLetterRunes = append(otherLetterRunes, r)
		}
	}
	for i, j := 0, len(englishLetters)>>1; i < j; i++ {
		englishLetters[i], englishLetters[len(englishLetters)-1-i] = englishLetters[len(englishLetters)-1-i], englishLetters[i]
	}
	result := make([]rune, len(rs))
	for i, idx := range otherLetterIndexes {
		result[idx] = otherLetterRunes[i]
	}
	for i, j := 0, 0; i < len(rs); i++ {
		if result[i] != 0 {
			continue
		}
		result[i] = englishLetters[j]
		j++
	}
	return string(result)
}

func main() {
	fmt.Println(reverseOnlyLetters("ab-cd"))
}
