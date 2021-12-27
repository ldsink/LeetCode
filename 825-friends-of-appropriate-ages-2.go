package main

func numFriendRequests(ages []int) int {
	userCount := make([]int, 121)
	accumulate := make([]int, 121)
	for _, age := range ages {
		userCount[age]++
	}
	for i, j := 1, 0; i < 121; i++ {
		j += userCount[i]
		accumulate[i] = j
	}

	var result int
	for i := 15; i < 121; i++ {
		if userCount[i] == 0 {
			continue
		}
		// 每个相同年龄的人之间的请求次数。
		result += userCount[i] * (userCount[i] - 1)
		// 不同年龄层之间的请求次数。
		start := i/2 + 8
		// 满足条件的全部人数
		sum := accumulate[i-1] - accumulate[start-1]
		result += userCount[i] * sum
	}
	return result
}
