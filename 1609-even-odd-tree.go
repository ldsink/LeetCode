package main

func isEvenOddTree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	evenNodes := []*TreeNode{root}
	oddNodes := []*TreeNode{}

	for len(oddNodes) > 0 || len(evenNodes) > 0 {
		// 检查偶数层
		for i := 0; i < len(evenNodes); i++ {
			if evenNodes[i].Val%2 == 0 {
				return false
			}
			if evenNodes[i].Left != nil {
				oddNodes = append(oddNodes, evenNodes[i].Left)
			}
			if evenNodes[i].Right != nil {
				oddNodes = append(oddNodes, evenNodes[i].Right)
			}
		}
		// 偶数层严格递减
		for i := 1; i < len(evenNodes); i++ {
			if evenNodes[i].Val <= evenNodes[i-1].Val {
				return false
			}
		}
		// 清空偶数层
		evenNodes = []*TreeNode{}

		// 检查奇数层
		for i := 0; i < len(oddNodes); i++ {
			if oddNodes[i].Val%2 != 0 {
				return false
			}
			if oddNodes[i].Left != nil {
				evenNodes = append(evenNodes, oddNodes[i].Left)
			}
			if oddNodes[i].Right != nil {
				evenNodes = append(evenNodes, oddNodes[i].Right)
			}
		}
		// 奇数层严格递减
		for i := 1; i < len(oddNodes); i++ {
			if oddNodes[i].Val >= oddNodes[i-1].Val {
				return false
			}
		}
		// 清空奇数层
		oddNodes = []*TreeNode{}
	}
	return true
}
