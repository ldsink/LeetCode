package main

import "strconv"

func atMostNGivenDigitSet(digits []string, n int) int {
	nStr := strconv.Itoa(n)
	fullCount := []int{1} // 每一位上的全部可能性
	for i := 0; i < len(nStr); i++ {
		fullCount = append(fullCount, fullCount[i]*len(digits))
	}
	var result int
	for i, success := 0, true; i < len(nStr); i++ {
		l := len(nStr) - i // 当前处理的长度
		if l > 1 {         // 大于 1 位的时候，0 开头肯定满足全部条件
			result += fullCount[l-1]
		}
		if !success { // 前缀已经失败，统计完当前长度的全排列就行
			continue
		}
		var j int
		for j = 0; j < len(digits); j++ {
			digit := digits[j][0]
			if digit < nStr[i] {
				result += fullCount[l-1]
			} else {
				if digit > nStr[i] { // 如果不相等，后面不会再有合法排列了
					success = false
				} else if l == 1 { // 最后一位的时候，相等也增加一种情况
					result++
				}
				break
			}
		}
		if success && j == len(digits) { // 最大值也处理到了，后面没有有效的解了
			success = false
		}
	}
	return result
}
