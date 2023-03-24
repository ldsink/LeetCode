package main

type StreamChecker struct {
	words   [][]rune
	letters []rune
}

func Constructor(words []string) StreamChecker {
	var w [][]rune
	for _, word := range words {
		w = append(w, []rune(word))
	}
	return StreamChecker{words: w}
}

func (this *StreamChecker) Query(letter byte) bool {
	this.letters = append(this.letters, rune(letter))
	if len(this.letters) > 2000 {
		this.letters = this.letters[len(this.letters)-200:]
	}
	var i, j int
	for _, word := range this.words {
		if len(word) > len(this.letters) {
			continue
		}
		for i, j = 0, len(this.letters)-len(word); i < len(word); i++ {
			if word[i] != this.letters[j+i] {
				break
			}
		}
		if i == len(word) {
			return true
		}
	}
	return false
}

/**
 * Your StreamChecker object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.Query(letter);
 */
