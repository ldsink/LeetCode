package main

type WordDictionary struct {
	children [26]*WordDictionary
	end      bool
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{}
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	if len(word) == 0 {
		this.end = true
		return
	}
	child := int(word[0] - 'a')
	if this.children[child] == nil {
		this.children[child] = &WordDictionary{}
	}
	this.children[child].AddWord(word[1:])
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	if len(word) == 0 {
		return this.end
	}
	if word[0] == '.' {
		for i := 0; i < 26; i++ {
			if this.children[i] != nil && this.children[i].Search(word[1:]) {
				return true
			}
		}
		return false
	}
	child := int(word[0] - 'a')
	if this.children[child] == nil {
		return false
	}
	return this.children[child].Search(word[1:])
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
