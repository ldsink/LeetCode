package main

func connect(root *Node) *Node {
	list := make([]*Node, 4096)
	list[0] = root
	for i := 0; list[i] != nil; i++ {
		if list[i].Left == nil {
			continue
		}
		right := (i + 1) << 1
		left := right - 1
		list[left] = list[i].Left
		list[right] = list[i].Right
	}
	for i := 1; list[(1<<uint(i))-1] != nil; i++ {
		j := 1<<uint(i) - 1
		k := 1<<uint(i+1) - 3
		for ; j <= k; j++ {
			list[j].Next = list[j+1]
		}
	}
	return root
}
