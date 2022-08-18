package main

func maxEqualFreq(nums []int) int {
	numCount, lenCount := make(map[int]int), make(map[int]int)
	for _, num := range nums {
		numCount[num]++
	}
	for _, count := range numCount {
		lenCount[count]++
	}

	deleteNum := func(num int) {
		lenCount[numCount[num]]--
		if lenCount[numCount[num]] == 0 {
			delete(lenCount, numCount[num])
		}
		numCount[num]--
		if numCount[num] == 0 {
			delete(numCount, num)
		} else {
			lenCount[numCount[num]]++
		}
	}

	isMatch := func() bool {
		if len(lenCount) > 2 { // 长度种类超过 2 肯定不行
			return false
		} else if len(lenCount) == 1 { // 数字的长度都相同，要么只有一个数，要么每个长度必须是 1 才符合条件
			if len(numCount) == 1 { // 只有一个数，符合条件
				return true
			}
			for l, _ := range lenCount {
				return l == 1 // 多个数，但是每个数的长度都是 1，符合条件
			}
		}
		// 长度类型只有 2 个，合法的场景有如下两种
		// 1. 长度相差 1，且长的长度的数量是 1
		// 2. 长度差别很大，但是短的长度为1 且只有 1 个数
		var data [][2]int
		for l, c := range lenCount {
			data = append(data, [2]int{l, c})
		}
		if data[0][0] > data[1][0] {
			data[0], data[1] = data[1], data[0] // 始终保持 data[0] 的长度小于 data[1]
		}
		if data[0][0]+1 == data[1][0] && data[1][1] == 1 { // 长度相差 1，且长的长度的数量是 1
			return true
		} else if data[0][0] == 1 && data[0][1] == 1 { // 长度差别很大，但是短的长度为1 且只有 1 个数
			return true
		}
		return false
	}

	for i := len(nums) - 1; i > 0; i-- {
		if isMatch() {
			return i + 1
		}
		deleteNum(nums[i])
	}
	return 1
}
