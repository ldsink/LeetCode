package main

func maxDistance(colors []int) int {
	if colors[0] != colors[len(colors)-1] {
		return len(colors) - 1
	}
	for i, j, k := 1, len(colors)>>1, len(colors)-1; i <= j; i++ {
		if colors[0] != colors[k-i] || colors[i] != colors[k] {
			return k - i
		}
	}
	return 0 // 上面一定会退出，这里只是个兜底的
}
