package main

type Trie struct {
	children [26]*Trie
	end      bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	if len(word) == 0 {
		this.end = true
		return
	}
	child := int(word[0] - 'a')
	if this.children[child] == nil {
		this.children[child] = &Trie{}
	}
	this.children[child].Insert(word[1:])
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	if len(word) == 0 {
		return this.end
	}
	child := int(word[0] - 'a')
	if this.children[child] == nil {
		return false
	}
	return this.children[child].Search(word[1:])
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	if len(prefix) == 0 {
		return true
	}
	child := int(prefix[0] - 'a')
	if this.children[child] == nil {
		return false
	}
	if len(prefix) == 1 {
		return true
	}
	return this.children[child].StartsWith(prefix[1:])

}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
