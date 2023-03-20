package main

import "strconv"

func numDupDigitsAtMostN(n int) int {
	multi := []int{1, 1, 2, 6, 24, 120, 720, 5040, 40320, 362880, 3628800}
	s := strconv.Itoa(n)
	var f func(int, int, int, bool) int
	f = func(used, count, position int, isPreEqual bool) int {
		if position == len(s) { // n 的所有位都不同时或者全 0 时
			return 1
		}
		max := 9
		if isPreEqual {
			max = int(s[position] - '0')
		}
		rst := 0
		for j := 0; j <= max; j++ { // 当前位取 j
			if used&(1<<j) != 0 { // j 已经被使用
				continue
			}
			if isPreEqual && j == max { // 之前都相等，这一位也相等，继续往下算，剩余数字个数减一
				rst += f(used|(1<<j), count-1, position+1, true)
			} else if used == 0 && j == 0 { // 还是前导 0 阶段，继续往下算，剩余数字个数不变（前导 0 不占据 0）
				rst += f(0, count, position+1, false)
			} else {
				// 不是前导 0，且到目前为止和 n 不完全相同，则之后的位顺便取
				// 剩余 len(s) - (position + 1) 位，可以使用 count - 1 个数字
				// 共 C(count - 1, len(s) - (position + 1)) * (len(s) - (position + 1))!
				//  = (count - 1)! / (count - len(s) + position)! 种情况
				t := count - len(s) + position // count = 10 - x，x <= position，count >= 10 - position，count - len(s) + position >= 10 - len(s) >= 0
				rst += multi[count-1] / multi[t]
			}
		}
		return rst
	}
	return n + 1 - f(0, 10, 0, true) // func 中计算了 0 的情况
}
