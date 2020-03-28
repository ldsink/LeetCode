package main

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p == nil {
		return false
	} else if q == nil {
		return false
	} else if (*q).Val != (*p).Val {
		return false
	}
	return isSameTree((*p).Left, (*q).Left) && isSameTree((*p).Right, (*q).Right)
}
