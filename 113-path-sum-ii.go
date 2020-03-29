package main

func pathSum(root *TreeNode, sum int) [][]int {
	var result [][]int
	if root == nil {
		return result
	} else if (*root).Left == nil && (*root).Right == nil {
		if (*root).Val == sum {
			result = append(result, []int{(*root).Val})
		}
		return result
	}
	if r := pathSum((*root).Left, sum-(*root).Val); len(r) > 0 {
		for _, path := range r {
			result = append(result, append([]int{(*root).Val}, path...))
		}
	}
	if r := pathSum((*root).Right, sum-(*root).Val); len(r) > 0 {
		for _, path := range r {
			result = append(result, append([]int{(*root).Val}, path...))
		}
	}
	return result
}
