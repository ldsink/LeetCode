package main

func maxPathSum(root *TreeNode) int {
	var nodes []*TreeNode
	nodes = append(nodes, root)
	var left, right []int
	for i := 0; i < len(nodes); i++ {
		left = append(left, 0)
		right = append(right, 0)
		if nodes[i].Left != nil {
			nodes = append(nodes, nodes[i].Left)
			left[i] = len(nodes) - 1
		}
		if nodes[i].Right != nil {
			nodes = append(nodes, nodes[i].Right)
			right[i] = len(nodes) - 1
		}
	}

	result := root.Val
	singlePath := make([]int, len(nodes))
	for i := len(nodes) - 1; i >= 0; i-- {
		singlePath[i] = nodes[i].Val
		val := nodes[i].Val
		if nodes[i].Left != nil && singlePath[left[i]] > 0 {
			val += singlePath[left[i]]
			if p := nodes[i].Val + singlePath[left[i]]; singlePath[i] < p {
				singlePath[i] = p
			}
		}
		if nodes[i].Right != nil && singlePath[right[i]] > 0 {
			val += singlePath[right[i]]
			if p := nodes[i].Val + singlePath[right[i]]; singlePath[i] < p {
				singlePath[i] = p
			}
		}
		if result < val {
			result = val
		}
	}
	return result
}
