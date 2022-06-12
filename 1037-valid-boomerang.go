package main

func isBoomerang(points [][]int) bool {
	type point struct {
		x, y int
	}
	a, b, c := point{x: points[0][0], y: points[0][1]}, point{x: points[1][0], y: points[1][1]}, point{x: points[2][0], y: points[2][1]}
	if a == b || b == c || c == a {
		return false
	}
	// 处理垂直 x 轴的情况
	if a.x == b.x {
		return c.x != a.x
	}
	k := float64(a.y-b.y) / float64(a.x-b.x)
	d := float64(a.y) - k*float64(a.x)

	isEqual := func(a, b float64) bool {
		if a < b {
			a, b = b, a
		}
		return a-b < 0.00000000001
	}
	return !isEqual(k*float64(c.x)+d, float64(c.y))
}
