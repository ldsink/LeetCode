package main

func divisibilityArray(word string, m int) []int {
	getNum := func(r rune) int {
		return int(r - '0')
	}
	var result []int
	var value int
	for _, r := range []rune(word) {
		value = value*10 + getNum(r)
		value %= m
		if value == 0 {
			result = append(result, 1)
		} else {
			result = append(result, 0)
		}
	}
	return result
}
