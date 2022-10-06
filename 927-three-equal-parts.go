package main

func threeEqualParts(arr []int) []int {
	var zeroOffset int
	for ; zeroOffset < len(arr); zeroOffset++ {
		if arr[zeroOffset] == 1 {
			break
		}
	}
	if zeroOffset == len(arr) { // 全部是 0 的情况
		return []int{0, 2}
	}

	// 此时从 zeroOffset 开始，需要将 arr 分成三段，如果成功，结果坐标都需要加上 zeroOffset。
	// 如果成功，假设每段长度为 l，结果一定为 [
	//    one(zeroOffset, zeroOffset+l),
	//    zero1(一段连续的 0),
	//    two[len(arr) - 2l, len(arr) - l]
	//    zero2(一段连续的 0),
	//    three[len(arr) - l, len(arr)]
	// ]
	l, zero1, zero2 := (len(arr)-zeroOffset)/3, 0, 0
	// 首先初始化 zero1 和 zero2，此时不包含长度为 l 的情况
	for i := zeroOffset + l + 1; i < len(arr); i++ {
		if arr[i] == 1 {
			break
		}
		zero1++
	}
	for i := len(arr) - l - 2; i > zeroOffset; i-- {
		if arr[i] == 1 {
			break
		}
		zero2++
	}
	impossible := []int{-1, -1}
	for ; l > 0; l-- {
		// 进入后先更新当长度为 l 的情况下，两段 0 区间的情况
		if arr[zeroOffset+l] == 0 {
			zero1++
		} else {
			zero1 = 0
		}
		if arr[len(arr)-l-1] == 0 {
			zero2++
		} else {
			zero2 = 0
		}
		if zeroOffset+3*l+zero1+zero2 < len(arr) { // 排除掉不可能的情况
			return impossible
		}
		if arr[len(arr)-l] == 0 { // 第三段有效位置不满足条件，前往下一个 l 进行计算
			continue
		}
		// 基于两段 0 区间进行剪枝搜索
		if part2End, part3Start := zeroOffset+2*l+zero1, len(arr)-l; part2End > part3Start || part2End+zero2 < part3Start {
			continue
		} else {
			success, part2Start := true, zeroOffset+l+zero1
			for i := 0; success && i < l; i++ {
				if arr[zeroOffset+i] == arr[part2Start+i] && arr[zeroOffset+i] == arr[part3Start+i] {
					continue
				}
				success = false
				break
			}
			if success {
				return []int{zeroOffset + l - 1, part2Start + l}
			}
		}
	}
	return impossible
}
