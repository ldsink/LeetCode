package main

type TrieNode struct {
	children map[byte]*TrieNode
	end      bool
}

type StreamChecker struct {
	root    *TrieNode
	letters []byte
}

func addTrieNode(root *TrieNode, word string) {
	for node, i := root, len(word)-1; i >= 0; i-- {
		if node.children == nil {
			node.children = make(map[byte]*TrieNode)
		}
		if _, ok := node.children[word[i]]; !ok {
			node.children[word[i]] = &TrieNode{}
		}
		node = node.children[word[i]]
		if i == 0 {
			node.end = true
		}
	}
}

func Constructor(words []string) StreamChecker {
	root := &TrieNode{}
	for _, word := range words {
		addTrieNode(root, word)
	}
	return StreamChecker{root: root}
}

func (this *StreamChecker) Query(letter byte) bool {
	this.letters = append(this.letters, letter)
	if len(this.letters) > 2000 {
		this.letters = this.letters[len(this.letters)-200:]
	}
	i := len(this.letters) - 1
	for node := this.root; i >= 0; i-- {
		if _, ok := node.children[this.letters[i]]; !ok {
			return false
		}
		node = node.children[this.letters[i]]
		if node.end {
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
