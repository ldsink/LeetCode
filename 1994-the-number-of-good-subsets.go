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
	mod := 1000000007
	var result int
	for i, n := 1, 1<<18; i < n; i++ {
		flag := 1  //flag在第一次进入时，成为now[j]，之后就乘以now[j];
		k := 1     //统计这种状态能形成的组合个数，因为乘法，所以初始化为1
		ok := true //是否合法的标记
		for j := 17; j >= 0; j-- {
			if (1<<j)&i > 0 {
				if flag == -1 {
					flag = validNums[j]
					k *= count[validNums[j]]
				} else {
					if gcd(flag, validNums[j]) != 1 {
						ok = false
						break
					}
					flag *= validNums[j]
					k *= count[validNums[j]]
					k %= mod
				}
			}
		}
		if ok {
			result = (result + k) % mod
		}
	}
	for i := 0; i < count[1]; i++ {
		result = (result * 2) % mod
	}
	return result
}
