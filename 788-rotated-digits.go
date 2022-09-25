package main

import (
	"strconv"
)

func rotatedDigits(n int) int {
	diffCount, diffSumCount, sameCount := []int{4}, []int{4}, []int{2}
	for i := 1; i <= 5; i++ {
		diffCount = append(diffCount, 7*diffCount[i-1]+4*sameCount[i-1])    // 长度刚好为 i 位合法的情况
		diffSumCount = append(diffSumCount, diffSumCount[i-1]+diffCount[i]) // 长度不超过 i 位合法的情况
		sameCount = append(sameCount, 3*sameCount[i-1])                     // 长度刚好为 i 位，翻转后和翻转前相同的情况
	}

	isValid := func(val int) bool {
		return !(val == 3 || val == 4 || val == 7)
	}
	isSame := func(val int) bool {
		return val == 0 || val == 1 || val == 8
	}
	pow := func(a, b int) (r int) {
		r = 1
		for i := 0; i < b; i++ {
			r *= a
		}
		return
	}

	var result int
	s := strconv.Itoa(n)
	samePrefix := true // 前缀翻转后是相同的值
	for i := 0; i < len(s); i++ {
		num := int(s[i] - '0')
		if idx := len(s) - i - 2; idx >= 0 {
			for j := 0; j < num; j++ {
				if isValid(j) {
					result += diffSumCount[idx]
					if !samePrefix || !isSame(j) { // 前缀翻转不相同，后面还可以多加上翻转相同的情况
						result += pow(3, idx+1)
					}
				}
			}
			if !isValid(num) {
				break
			} else if samePrefix && !isSame(num) {
				samePrefix = false
			}
		} else { // idx < 0，已经到最后一位了
			for j := 0; j <= num; j++ {
				if isValid(j) && (!samePrefix || !isSame(j)) {
					result++
				}
			}
		}
	}
	return result
}
