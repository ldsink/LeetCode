package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	// 如果 depth == 1 意味着 depth - 1 根本没有深度，那么创建一个树节点，值 val 作为整个原始树的新根，而原始树就是新根的左子树。
	if depth == 1 {
		return &TreeNode{Val: val, Left: root}
	} else if depth == 2 {
		// 给定整数 depth，对于深度为 depth - 1 的每个非空树节点 cur ，创建两个值为 val 的树节点作为 cur 的左子树根和右子树根。
		root.Left = &TreeNode{Val: val, Left: root.Left}
		root.Right = &TreeNode{Val: val, Right: root.Right}
		return root
	}
	if root.Left != nil {
		root.Left = addOneRow(root.Left, val, depth-1)
	}
	if root.Right != nil {
		root.Right = addOneRow(root.Right, val, depth-1)
	}
	return root
}
