package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	nodes := []*TreeNode{root}
	var s []string
	for i := 0; i < len(nodes); i++ {
		var ln, rn int
		if c := nodes[i].Left; c != nil {
			nodes = append(nodes, c)
			ln = 1
		}
		if c := nodes[i].Right; c != nil {
			nodes = append(nodes, c)
			rn = 1
		}
		s = append(s, fmt.Sprintf("%d:%d:%d", nodes[i].Val, ln, rn))
	}
	return strings.Join(s, ";")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	s := strings.Split(data, ";")
	parentNodes := []*TreeNode{nil}
	isLeft := make(map[int]bool)
	for i, d := range s {
		n := strings.SplitN(d, ":", 3)

		val, _ := strconv.Atoi(n[0])
		node := TreeNode{Val: val}
		if parentNodes[i] == nil { // root node，set parent to itself
			parentNodes[i] = &node
		} else if isLeft[i] {
			parentNodes[i].Left = &node
		} else {
			parentNodes[i].Right = &node
		}

		hasLeft, _ := strconv.Atoi(n[1])
		if hasLeft == 1 {
			parentNodes = append(parentNodes, &node)
			isLeft[len(parentNodes)-1] = true
		}
		hasRight, _ := strconv.Atoi(n[2])
		if hasRight == 1 {
			parentNodes = append(parentNodes, &node)
		}
	}
	return parentNodes[0]
}
