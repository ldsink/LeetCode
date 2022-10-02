package main

func canTransform(start string, end string) bool {
	const L = 'L'
	const X = 'X'
	const R = 'R'
	// L 和 R 数量和顺序相同
	var src, dst []rune
	for _, r := range []rune(start) {
		if r != X {
			src = append(src, r)
		}
	}
	for _, r := range []rune(end) {
		if r != X {
			dst = append(dst, r)
		}
	}
	if string(src) != string(dst) {
		return false
	}
	// src 的 L 前面的 X 不能大于 dst
	var count, index int
	var xCount []int
	for _, r := range []rune(end) {
		if r == X {
			count++
		} else if r == L {
			xCount = append(xCount, count)
		}
	}
	count = 0
	for _, r := range []rune(start) {
		if r == X {
			count++
		} else if r == L {
			if xCount[index] > count {
				return false
			}
			index++
		}
	}
	// src 的 R 后面的 X 不能大于 dst
	count, index, xCount = 0, 0, []int{}
	for i := len(end) - 1; i >= 0; i-- {
		if end[i] == X {
			count++
		} else if end[i] == R {
			xCount = append(xCount, count)
		}
	}
	count = 0
	for i := len(start) - 1; i >= 0; i-- {
		if start[i] == X {
			count++
		} else if start[i] == R {
			if xCount[index] > count {
				return false
			}
			index++
		}
	}
	return true
}
