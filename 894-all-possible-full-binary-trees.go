package main

func allPossibleFBT(n int) []*TreeNode {
	if n%2 == 0 {
		return []*TreeNode{}
	}
	fbt := make(map[int][]*TreeNode)
	fbt[1] = []*TreeNode{{Val: 0}}
	fbt[3] = []*TreeNode{{Val: 0, Left: fbt[1][0], Right: fbt[1][0]}}
	for i := 5; i <= n; i += 2 {
		for j := 1; j+1 < i; j += 2 {
			k := i - 1 - j
			for _, left := range fbt[j] {
				for _, right := range fbt[k] {
					fbt[i] = append(fbt[i], &TreeNode{Val: 0, Left: left, Right: right})
				}
			}
		}
	}
	return fbt[n]
}
