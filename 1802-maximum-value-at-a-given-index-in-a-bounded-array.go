package main

func maxValue(n int, index int, maxSum int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	getSum := func(left, right, val int) int {
		s := val
		if left > 0 {
			s += ((val - 1 + val - left) * left) >> 1
		}
		if right > 0 {
			s += ((val - 1 + val - right) * right) >> 1
		}
		return s
	}
	getVal := func(left, right, val int) int {
		sum := val
		if val >= left {
			start, end := val-left, val-1
			sum += int(float64(start+end) / 2 * float64(left))
		} else {
			sum += (val * (val - 1)) >> 1
		}
		if val >= right {
			start, end := val-right, val-1
			sum += int(float64(start+end) / 2 * float64(right))
		} else {
			sum += (val * (val - 1)) >> 1
		}
		return sum
	}
	left := index
	right := n - index - 1
	val := max(left, right) + 1
	// index 位置在取最大值的情况下，当前的总和
	s := getSum(left, right, val)
	// 根据和 maxSum 的比较，整体上移或者下移
	if s <= maxSum {
		val += (maxSum - s) / n
	} else {
		maxSum -= n
		var start, end int
		for start, end = 0, max(left, right); start < end; {
			middle := (start + end) >> 1
			if getVal(left, right, middle+1) > maxSum {
				end = middle
			} else {
				start = middle + 1
			}
		}
		val = start + 1
	}
	return val
}
