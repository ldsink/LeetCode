package main

func sumRootToLeaf(root *TreeNode) int {
	var result int
	var search func(*TreeNode, int)
	search = func(node *TreeNode, num int) {
		num = (num << 1) | node.Val
		if node.Left == nil && node.Right == nil {
			result += num
			return
		}
		if node.Left != nil {
			search(node.Left, num)
		}
		if node.Right != nil {
			search(node.Right, num)
		}
	}
	search(root, 0)
	return result
}
