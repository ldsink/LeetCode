package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMiddle(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Left == nil {
		return root
	}
	return findMiddle(root.Left)
}

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	if p.Val < root.Val {
		result := inorderSuccessor(root.Left, p)
		if result == nil {
			result = root
		}
		return result
	} else if root.Val == p.Val {
		return findMiddle(root.Right)
	}
	return inorderSuccessor(root.Right, p)
}
