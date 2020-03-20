package main

func trap(height []int) int {
	var result int
	var blocks [][3]int // top, bottom, index
	for i, h := range height {
		var newBlocks [][3]int
		for _, v := range blocks {
			top, bottom, index := v[0], v[1], v[2]
			if bottom >= h {
				newBlocks = append(newBlocks, [3]int{top, bottom, index})
			} else if bottom <= h && h <= top {
				result += (i - index - 1) * (h - bottom)
				if top > h {
					newBlocks = append(newBlocks, [3]int{top, h, index})
				}
			} else if top <= h {
				result += (i - index - 1) * (top - bottom)
			}
		}
		newBlocks = append(newBlocks, [3]int{h, 0, i})
		blocks = newBlocks
	}
	return result
}
