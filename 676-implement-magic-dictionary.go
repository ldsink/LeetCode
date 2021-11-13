package main

type MagicDictionary struct {
	dictionary []string
}

func Constructor() MagicDictionary {
	return MagicDictionary{}
}

func (this *MagicDictionary) BuildDict(dictionary []string) {
	this.dictionary = dictionary
}

func (this *MagicDictionary) Search(searchWord string) bool {
	var diff int
	for i := 0; i < len(this.dictionary); i++ {
		if len(this.dictionary[i]) != len(searchWord) {
			continue
		}
		diff = 0
		for j := 0; j < len(searchWord) && diff < 2; j++ {
			if searchWord[j] != this.dictionary[i][j] {
				diff++
			}
		}
		if diff == 1 {
			return true
		}
	}
	return false
}

/**
 * Your MagicDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.BuildDict(dictionary);
 * param_2 := obj.Search(searchWord);
 */
