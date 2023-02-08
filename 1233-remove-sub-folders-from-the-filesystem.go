package main

import (
	"sort"
	"strings"
)

type trieNode struct {
	children map[string]*trieNode
	exist    bool
}

type trieTree struct {
	root *trieNode
}

func (t *trieTree) addPath(path string) bool {
	if path == "" || path == "/" {
		return false
	}
	parts := strings.Split(path, "/")
	for i, current := 0, t.root; i <= len(parts); i++ {
		if i == len(parts) {
			current.exist = true
			return true
		}
		if current.exist {
			return false
		}
		if current.children == nil {
			current.children = make(map[string]*trieNode)
		}
		if current.children[parts[i]] == nil {
			current.children[parts[i]] = &trieNode{}
		}
		current = current.children[parts[i]]
	}
	return true
}

func removeSubfolders(folder []string) []string {
	sort.Strings(folder)
	tree := &trieTree{root: &trieNode{}}
	var result []string
	for _, path := range folder {
		if tree.addPath(path) {
			result = append(result, path)
		}
	}
	return result
}
