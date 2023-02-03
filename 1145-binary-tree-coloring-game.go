package main

func btreeGameWinningMove(root *TreeNode, n int, x int) bool {
	var xNode *TreeNode
	var upper, left, right int
	var findX func(*TreeNode)
	findX = func(node *TreeNode) {
		if node.Val == x {
			xNode = node
			return
		}
		upper++
		if node.Left != nil {
			findX(node.Left)
		}
		if node.Right != nil {
			findX(node.Right)
		}
	}
	var countNode func(*TreeNode) int
	countNode = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		return 1 + countNode(node.Left) + countNode(node.Right)
	}
	findX(root)
	left = countNode(xNode.Left)
	right = countNode(xNode.Right)
	n >>= 1
	return upper > n || left > n || right > n
}
