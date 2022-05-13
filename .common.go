package main

/*
	公共数据定义和常见的函数，在写题目的时候可以快速复制进去。
	由于名称以'.'开头，'.common.go' 被构建工具忽略
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func abs(a int) {
	if a >= 0 {
		return a
	}
	return -a
}
