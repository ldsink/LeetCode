package main

func uniqueLetterString(s string) int {
	const A = 'A'
	var indexes [26][]int
	for i, r := range s {
		indexes[r-A] = append(indexes[r-A], i)
	}

	countByIndex := func(index []int) int {
		if len(index) == 0 { // 没有出现这个字符
			return 0
		}
		var count int
		for i := 0; i < len(index); i++ {
			prev, current, next := 0, index[i], len(s)
			if i != 0 { // 获得包含当前位置的起点的坐标
				prev = index[i-1] + 1
			}
			if i != len(index)-1 { // 获得包含当前位置的终点的坐标
				next = index[i+1]
			}
			count += (current - prev + 1) * (next - current)
		}
		return count
	}

	var result int
	for i := 0; i < 26; i++ {
		result += countByIndex(indexes[i])
	}
	return result
}
