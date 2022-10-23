package main

func mergeAlternately(word1 string, word2 string) string {
	max := len(word1)
	var extend string
	if len(word1) > len(word2) {
		max = len(word2)
		extend = word1[max:]
	} else {
		extend = word2[max:]
	}
	var result []rune
	for i := 0; i < max; i++ {
		result = append(result, rune(word1[i]))
		result = append(result, rune(word2[i]))
	}
	result = append(result, []rune(extend)...)
	return string(result)
}
