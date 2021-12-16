package main

import (
	"math"
	"sort"
)

func visiblePoints(points [][]int, angle int, location []int) int {
	// 不与 location 位置重合的点的角度
	var angles []float64
	// 和 location 重合的点的数量
	var same int
	for i := 0; i < len(points); i++ {
		x, y := points[i][0]-location[0], points[i][1]-location[1]
		// 重合的情况
		if x == 0 && y == 0 {
			same++
			continue
		}
		// 0°，90°，180° 和 270° 的情况
		if y == 0 {
			if x > 0 {
				angles = append(angles, 0)
			} else {
				angles = append(angles, 180)
			}
		} else if x == 0 {
			if y > 0 {
				angles = append(angles, 90)
			} else {
				angles = append(angles, 270)
			}
		} else if x > 0 && y > 0 {
			// 处理第一象限
			tan := math.Abs(float64(y) / float64(x))
			degree := math.Atan(tan) * 180 / math.Pi
			angles = append(angles, degree)
		} else if x < 0 && y > 0 {
			// 处理第二象限
			tan := math.Abs(float64(y) / float64(x))
			degree := math.Atan(tan) * 180 / math.Pi
			degree = 180 - degree
			angles = append(angles, degree)
		} else if x < 0 && y < 0 {
			// 处理第三象限
			tan := math.Abs(float64(y) / float64(x))
			degree := math.Atan(tan) * 180 / math.Pi
			degree = 180 + degree
			angles = append(angles, degree)
		} else {
			// 处理第四象限
			tan := math.Abs(float64(y) / float64(x))
			degree := math.Atan(tan) * 180 / math.Pi
			degree = 360 - degree
			angles = append(angles, degree)
		}
	}

	// 没有外面的点或者只有一个点了
	if len(angles) <= 1 {
		return same + len(angles)
	}

	sort.Float64s(angles)
	var max int // 在 angle 角度内能看到的最多的点的数量
	// start 开始点的索引，从这个点往逆时针方向看 angle 度
	for start, count := 0, 2; start < len(angles); start++ {
		// 把上个点移除掉，然后确保包含自己
		if count > 1 {
			count--
		}
		// 扩展范围
		for ; count < len(angles); count++ {
			idx := (start + count + len(angles)) % len(angles)
			diff := angles[idx] - angles[start]
			if diff < 0 {
				diff += 360
			}
			if diff > float64(angle) {
				break
			}
		}
		// 可以看到全部的点，停止循环
		if count == len(angles) {
			max = count
			break
		}
		if max < count {
			max = count
		}
	}
	return same + max
}
