package main

func isCousins(root *TreeNode, x int, y int) bool {
	type Node struct {
		node   *TreeNode
		parent *TreeNode
		depth  int
	}
	getDepthAndParent := func(v int) (int, *TreeNode) {
		queue := []Node{{node: root, depth: 0}}
		for i := 0; i < len(queue); i++ {
			if queue[i].node.Val == v {
				return queue[i].depth, queue[i].parent
			}
			if queue[i].node.Left != nil {
				queue = append(queue, Node{node: queue[i].node.Left, parent: queue[i].node, depth: queue[i].depth + 1})
			}
			if queue[i].node.Right != nil {
				queue = append(queue, Node{node: queue[i].node.Right, parent: queue[i].node, depth: queue[i].depth + 1})
			}
		}
		return -1, nil
	}
	xDepth, xParent := getDepthAndParent(x)
	yDepth, yParent := getDepthAndParent(y)
	return xDepth == yDepth && xParent != yParent
}
