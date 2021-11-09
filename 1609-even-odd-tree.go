package main

func isEvenOddTree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	nodes := []*TreeNode{root}
	nodeLevels := []int{0}
	var levelVal []int
	for i := 0; i < len(nodes); i++ {
		// 检查当前节点
		isOdd := nodeLevels[i]%2 != 0
		if isOdd {
			if nodes[i].Val%2 != 0 {
				return false
			}
		} else {
			if nodes[i].Val%2 == 0 {
				return false
			}
		}
		// 比较当前层级的值
		if len(levelVal) < nodeLevels[i]+1 {
			// 当前层级的第一个节点
			levelVal = append(levelVal, nodes[i].Val)
		} else {
			// 检查是不是按照顺序
			if isOdd {
				if levelVal[nodeLevels[i]] <= nodes[i].Val {
					return false
				}
			} else {
				if levelVal[nodeLevels[i]] >= nodes[i].Val {
					return false
				}
			}
			levelVal[nodeLevels[i]] = nodes[i].Val
		}
		// 添加下一层
		if nodes[i].Left != nil {
			nodes = append(nodes, nodes[i].Left)
			nodeLevels = append(nodeLevels, nodeLevels[i]+1)
		}
		if nodes[i].Right != nil {
			nodes = append(nodes, nodes[i].Right)
			nodeLevels = append(nodeLevels, nodeLevels[i]+1)
		}
	}
	return true
}
