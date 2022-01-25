package main

type DetectSquares struct {
	vertical map[int][]int
	points   map[[2]int]int
}

func Constructor() DetectSquares {
	return DetectSquares{vertical: make(map[int][]int), points: make(map[[2]int]int)}
}

func (this *DetectSquares) Add(point []int) {
	x, y := point[0], point[1]
	this.points[[2]int{x, y}]++
	if this.points[[2]int{x, y}] == 1 { // 只在第一次添加
		this.vertical[x] = append(this.vertical[x], y)
	}
}

func (this *DetectSquares) getEdge(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func (this *DetectSquares) Count(point []int) int {
	var count int
	x, y := point[0], point[1]
	for _, y2 := range this.vertical[x] {
		if y == y2 {
			continue
		}
		edge := this.getEdge(y, y2)
		// 左边
		count += this.points[[2]int{x - edge, y}] * this.points[[2]int{x - edge, y2}] * this.points[[2]int{x, y2}]
		// 右边
		count += this.points[[2]int{x + edge, y}] * this.points[[2]int{x + edge, y2}] * this.points[[2]int{x, y2}]
	}
	return count
}

/**
 * Your DetectSquares object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(point);
 * param_2 := obj.Count(point);
 */
