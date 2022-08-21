package main

func isPrefixOfWord(sentence string, searchWord string) int {
	var start, end int

	isPrefix := func() bool {
		if end-start < len(searchWord) {
			return false
		}
		return sentence[start:start+len(searchWord)] == searchWord
	}

	for idx := 1; end < len(sentence); start, end, idx = end+1, end+1, idx+1 {
		for ; end < len(sentence) && sentence[end] != ' '; end++ {
		}
		if isPrefix() {
			return idx
		}
	}
	return -1
}
