package main

func flipLights(n int, presses int) int {
	// bits for odd, even, oddLabel, evenLabel conditions
	conditions := 1
	if n >= 4 {
		conditions = 4
	} else if n >= 3 {
		conditions = 3
	} else if n >= 2 {
		conditions = 2
	}
	initState, mask := (1<<conditions)-1, (1<<conditions)-1

	var result int
	searched := make(map[int16]bool) // 6 bit for state, 10bit for presses
	exists := make([]bool, initState+1)
	var search func(int, int)
	search = func(state int, count int) {
		if count == 0 {
			if !exists[state] {
				exists[state] = true
				result++
			}
			return
		}

		key := int16(state<<10 | count)
		if searched[key] {
			return
		}
		searched[key] = true

		// 反转当前所有灯的状态（即开变为关，关变为开）
		newState := ^state
		search(newState&mask, count-1)
		// 反转编号为奇数的灯的状态（即 1, 3, ...）
		newState = state ^ (1 << 0) ^ (1 << 2)
		search(newState&mask, count-1)
		// 反转编号为偶数的灯的状态（即 2, 4, ...）
		newState = state ^ (1 << 1) ^ (1 << 3)
		search(newState&mask, count-1)
		// 反转编号为 j = 3k + 1 的灯的状态，其中 k = 0, 1, 2, ...（即 1, 4, 7, 10, ...）
		if n > 2 { // 从 3 个元素起这种情况才和简单奇偶有区别
			newState = state ^ (1 << 2) ^ (1 << 3)
			search(newState&mask, count-1)
		}
	}
	search(initState, presses)
	return result
}
