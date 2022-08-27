package main

func widthOfBinaryTree(root *TreeNode) int {
	var curr, prev []*TreeNode
	root.Val = 0 // 使用 Val 标记节点在当前层的 Index
	curr = []*TreeNode{root}
	var result int
	for len(curr) > 0 {
		// 判断当前层是否可以更新最大宽度
		if l := curr[len(curr)-1].Val - curr[0].Val + 1; result < l {
			result = l
		}
		// 开始扩展下一层
		prev, curr = curr, []*TreeNode{}
		for _, node := range prev {
			if node.Left != nil {
				node.Left.Val = node.Val << 1
				curr = append(curr, node.Left)
			}
			if node.Right != nil {
				node.Right.Val = node.Val<<1 + 1
				curr = append(curr, node.Right)
			}
		}
	}
	return result
}
