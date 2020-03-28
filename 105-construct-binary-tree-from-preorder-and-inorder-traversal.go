package main

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	} else if len(preorder) == 1 {
		return &TreeNode{Val: preorder[0]}
	}
	num := preorder[0]
	var count int
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == num {
			count = i
			break
		}
	}
	leftPreorder := preorder[1 : 1+count]
	rightPreorder := preorder[1+count:]
	leftInorder := inorder[0:count]
	rightInorder := inorder[count+1:]
	node := TreeNode{Val: num, Left: buildTree(leftPreorder, leftInorder), Right: buildTree(rightPreorder, rightInorder)}
	return &node
}
