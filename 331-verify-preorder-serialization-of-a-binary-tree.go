package main

import (
	"strings"
)

func isValidSerialization(preorder string) bool {
	nodes := strings.Split(preorder, ",")
	var leaf, node int
	for i := len(nodes) - 1; i >= 0; i-- {
		if nodes[i] == "#" {
			leaf++
		} else {
			node++
		}
		if leaf <= node {
			return false
		}
	}
	return node+1 == leaf
}
