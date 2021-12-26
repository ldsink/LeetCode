package main

func numFriendRequests(ages []int) int {
	userCount := make([]int, 121)
	for _, age := range ages {
		userCount[age]++
	}
	var distinctAges []int
	for age := 1; age < len(userCount); age++ {
		if userCount[age] > 0 {
			distinctAges = append(distinctAges, age)
		}
	}

	var result int
	// 先计算不同年龄的人的好友请求次数
	for i := 1; i < len(distinctAges); i++ {
		// age[y] > age[x] 不会发送请求，所以 j 必须小于 i
		for j := i - 1; j >= 0; j-- {
			// age[y] <= 0.5 * age[x] + 7 不会发送请求，所以更小的 j 可以忽略
			if float64(distinctAges[j]) <= 0.5*float64(distinctAges[i])+7 {
				break
			}
			// 这两个不同年龄层的用户可以加好友
			result += userCount[distinctAges[i]] * userCount[distinctAges[j]]
		}
	}
	// 每个相同年龄的人之间的请求次数。
	// 按照条件，age[y] <= 0.5 * age[x] + 7，必须大于 14 的同龄人才可以交朋友。
	for _, age := range distinctAges {
		if age <= 14 || userCount[age] <= 1 {
			continue
		}
		result += userCount[age] * (userCount[age] - 1)
	}
	return result
}
