package main

const PendingFlag = -1
const DrawFlag = 0
const MouseWinFlag = 1
const CatWinFlag = 2

func dp(cat, mouse, step int, searched [][][]int, graph [][]int) int {
	if mouse == 0 {
		return MouseWinFlag
	} else if cat == mouse {
		return CatWinFlag
	} else if searched[cat][mouse][step] != PendingFlag {
		return searched[cat][mouse][step]
	}
	var catWin, mouseWin, draw bool
	if step%2 == 1 {
		// 老鼠移动的情况，如果有一个可以达成胜利，那么就是胜利；如果所有的都是猫胜利，那么猫胜利
		for _, node := range graph[mouse] {
			r := dp(cat, node, step+1, searched, graph)
			if r == MouseWinFlag {
				mouseWin = true
				break
			} else if r == DrawFlag {
				draw = true
			}
		}
		if mouseWin {
			searched[cat][mouse][step] = MouseWinFlag
		} else if draw {
			searched[cat][mouse][step] = DrawFlag
		} else {
			searched[cat][mouse][step] = CatWinFlag
		}
	} else {
		// 猫移动的情况，如果有一个可以达成胜利，那么就是胜利；如果所有的都是老鼠胜利，那么老鼠胜利
		for _, node := range graph[cat] {
			if node == 0 { // 猫不可能进入洞中
				continue
			}
			r := dp(node, mouse, step+1, searched, graph)
			if r == CatWinFlag {
				catWin = true
				break
			} else if r == DrawFlag {
				draw = true
			}
		}
		if catWin {
			searched[cat][mouse][step] = CatWinFlag
		} else if draw {
			searched[cat][mouse][step] = DrawFlag
		} else {
			searched[cat][mouse][step] = MouseWinFlag
		}
	}
	return searched[cat][mouse][step]
}

func catMouseGame(graph [][]int) int {
	// 猫在的位置，老鼠在的位置，谁的回合。0 老鼠动，1 猫动
	var searched [][][]int
	for i := 0; i < len(graph); i++ {
		searched = append(searched, [][]int{})
		for j := 0; j < len(graph); j++ {
			var data []int
			for k := 0; k < len(graph)*2; k++ {
				data = append(data, PendingFlag)
			}
			data = append(data, DrawFlag)
			searched[i] = append(searched[i], data)
		}
	}
	// 动态规划
	return dp(2, 1, 1, searched, graph)
}
