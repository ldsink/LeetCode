package main

/**
 * Definition for a QuadTree node.
 * type Node struct {
 *     Val bool
 *     IsLeaf bool
 *     TopLeft *Node
 *     TopRight *Node
 *     BottomLeft *Node
 *     BottomRight *Node
 * }
 */

func construct(grid [][]int) *Node {
	initValue := grid[0][0]
	same := true
	for i := 0; same && i < len(grid); i++ {
		for j := 0; same && j < len(grid[i]); j++ {
			if initValue != grid[i][j] {
				same = false
			}
		}
	}
	if same {
		return &Node{Val: initValue == 1, IsLeaf: true}
	}
	// 拆分成四个子块
	var topLeft, topRight, bottomLeft, bottomRight [][]int
	middle := len(grid) / 2
	for i := 0; i < middle; i++ {
		topLeft = append(topLeft, grid[i][:middle])
		topRight = append(topRight, grid[i][middle:])
	}
	for i := middle; i < len(grid); i++ {
		bottomLeft = append(bottomLeft, grid[i][:middle])
		bottomRight = append(bottomRight, grid[i][middle:])
	}
	return &Node{
		Val:         true,
		IsLeaf:      false,
		TopLeft:     construct(topLeft),
		TopRight:    construct(topRight),
		BottomLeft:  construct(bottomLeft),
		BottomRight: construct(bottomRight),
	}
}
