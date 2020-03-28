package main

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	} else if len(inorder) == 1 {
		return &TreeNode{Val: inorder[0]}
	}
	num := postorder[len(postorder)-1]
	var count int
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == num {
			count = i
			break
		}
	}
	leftInorder := inorder[0:count]
	rightInorder := inorder[count+1:]
	leftPostorder := postorder[0:count]
	rightPostorder := postorder[count : len(postorder)-1]
	node := TreeNode{Val: num, Left: buildTree(leftInorder, leftPostorder), Right: buildTree(rightInorder, rightPostorder)}
	return &node
}
