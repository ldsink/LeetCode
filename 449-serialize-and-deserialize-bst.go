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
	// val,leftLength,leftData,rightLength,rightData
	leftData := this.serialize(root.Left)
	rightData := this.serialize(root.Right)
	return fmt.Sprintf("%d,%d,%s%s", root.Val, len(leftData), leftData, rightData)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}
	parts := strings.SplitN(data, ",", 3)
	val, _ := strconv.Atoi(parts[0])
	leftLength, _ := strconv.Atoi(parts[1])
	left := this.deserialize(parts[2][:leftLength])
	right := this.deserialize(parts[2][leftLength:])
	return &TreeNode{Val: val, Left: left, Right: right}
}
