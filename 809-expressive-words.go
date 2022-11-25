package main

func expressiveWords(s string, words []string) int {
	getLength := func(str string) int {
		for i := 1; i < len(str); i++ {
			if str[i] != str[0] {
				return i
			}
		}
		return len(str)
	}
	var isStretchy func(string, string) bool
	isStretchy = func(src, dst string) bool {
		if len(src) > len(dst) {
			return false
		} else if len(src) == len(dst) {
			return src == dst
		} else if len(src) == 0 || src[0] != dst[0] {
			return false
		}
		a, b := getLength(src), getLength(dst)
		if a > b {
			return false
		} else if a == b {
			return isStretchy(src[a:], dst[b:])
		} else if b < 3 {
			return false
		}
		return isStretchy(src[a:], dst[b:])
	}
	var count int
	for _, word := range words {
		if isStretchy(word, s) {
			count++
		}
	}
	return count
}
