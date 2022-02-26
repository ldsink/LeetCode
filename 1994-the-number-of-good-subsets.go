package main

func gcd(a, b int) int {
	if a%b == 0 {
		return b
	}
	return gcd(b, a%b)
}

func numberOfGoodSubsets(nums []int) int {
	validNums := []int{2, 3, 5, 6, 7, 10, 11, 13, 14, 15, 17, 19, 21, 22, 23, 26, 29, 30}
	count := make([]int, 31)
	for _, num := range nums {
		count[num]++
	}
	mask := 0 // mask 用来标记有效数字是否存在
	for idx, num := range validNums {
		if count[num] == 0 {
			mask ^= 1 << idx
		}
	}
	mod := 1000000007
	var result int
	for i, end := 1, 1<<18; i < end; i++ {
		if (0^i)&mask > 0 {
			continue // 当前组合包括不存在的数字，跳过
		}
		var currMulti, currCount int // 当前状态的乘积，当前的组合个数
		// 开始计算当前组合是否合法并且计算数量
		for idx, num := range validNums {
			if (1<<idx)&i == 0 {
				continue // 当前组合不包括这个数字，跳过
			}
			if currMulti == 0 {
				currMulti = num
				currCount = count[num]
			} else {
				if gcd(currMulti, num) != 1 {
					currMulti = 0 // 重新设置为 0，表示不合法，并且退出
					break
				}
				currMulti *= num
				currCount *= count[num]
				currCount %= mod
			}
		}
		if currMulti != 0 {
			result = (result + currCount) % mod
		}
	}
	// 每个位置上的 1 都可以采用采纳或者不采纳的方式，相当于方案 *2，进行处理
	for i := 0; i < count[1]; i++ {
		result = (result * 2) % mod
	}
	return result
}
