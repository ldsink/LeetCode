package main

func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}
	safeRob := make(map[*TreeNode]int)  // 可以抢这个点
	alertRob := make(map[*TreeNode]int) // 不能抢这个点
	stack := []*TreeNode{root}
	for i := 0; i < len(stack); i++ {
		if stack[i].Left != nil {
			stack = append(stack, stack[i].Left)
		}
		if stack[i].Right != nil {
			stack = append(stack, stack[i].Right)
		}
	}
	for i := len(stack) - 1; i >= 0; i-- {
		var takeIt, ignoreIt int
		if stack[i].Left != nil {
			ignoreIt += alertRob[stack[i].Left]
			takeIt += safeRob[stack[i].Left]
		}
		if stack[i].Right != nil {
			ignoreIt += alertRob[stack[i].Right]
			takeIt += safeRob[stack[i].Right]
		}
		// 当前节点不获取的最大金额
		safeRob[stack[i]] = ignoreIt
		// 当前节点获取的最大金额
		takeIt += stack[i].Val
		if takeIt < ignoreIt {
			takeIt = ignoreIt
		}
		alertRob[stack[i]] = takeIt
	}
	return alertRob[root]
}
