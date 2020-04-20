package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
	stack []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	iterator := BSTIterator{}
	for node := root; node != nil; {
		next := node.Left
		node.Left = nil
		iterator.stack = append(iterator.stack, node)
		node = next
	}
	return iterator
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	node := this.stack[len(this.stack)-1]
	num := node.Val
	this.stack = this.stack[:len(this.stack)-1]
	for node = node.Right; node != nil; {
		next := node.Left
		node.Left = nil
		this.stack = append(this.stack, node)
		node = next
	}
	return num
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return len(this.stack) > 0
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
