package main

func findClosestElements(arr []int, k int, x int) []int {
	if k == len(arr) {
		return arr
	} else if x <= arr[0] {
		return arr[:k]
	} else if x >= arr[len(arr)-1] {
		return arr[len(arr)-k:]
	}

	abs := func(a int) int {
		if a >= 0 {
			return a
		}
		return -a
	}

	// 两种简单处理策略:
	// 1. 如果 k 大于一半的数，那从两端做减法
	// 2. 否则找到接近 x 的位置，向两边扩展数值
	// 有时间还可以用二分找两边的点

	left, right := 0, len(arr)-1
	// 两边往中间更新
	if k >= len(arr)>>1 {
		for i := len(arr) - k; i > 0; i-- {
			if abs(arr[left]-x) <= abs(arr[right]-x) {
				right--
			} else {
				left++
			}
		}
		return arr[left : right+1]
	}
	// 中间往两边更新
	for left < right {
		middle := (left + right) >> 1
		if arr[middle] > x {
			right = middle
		} else {
			left = middle + 1
		}
	}
	left-- // 此时 left + 1 == right
	for ; k > 0 && left >= 0 && right < len(arr); k-- {
		if abs(arr[left]-x) <= abs(arr[right]-x) {
			left--
		} else {
			right++
		}
	}
	if k > 0 { // 还需要取数，但是一边空了
		if left >= 0 {
			left -= k
		} else {
			right += k
		}
	}
	return arr[left+1 : right]
}
