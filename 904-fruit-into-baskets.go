package main

func totalFruit(fruits []int) (result int) {
	zipFruits := func() (zippedFruits [][2]int) {
		curr := [2]int{fruits[0], 1}
		for i := 1; i < len(fruits); i++ {
			if curr[0] == fruits[i] {
				curr[1]++
			} else {
				zippedFruits = append(zippedFruits, curr)
				curr = [2]int{fruits[i], 1}
			}
		}
		zippedFruits = append(zippedFruits, curr)
		return
	}
	var buckets [][3]int
	inBucket := func(fruit [2]int) int {
		for i, bucket := range buckets {
			if bucket[0] == fruit[0] {
				return i
			}
		}
		return -1
	}
	getFruitsNum := func() (num int) {
		for _, bucket := range buckets {
			num += bucket[1]
		}
		return
	}
	for _, fruit := range zipFruits() {
		if i := inBucket(fruit); i >= 0 {
			buckets[i][1] += fruit[1]
			buckets[i][2] = fruit[1]
			if i == 0 && len(buckets) == 2 {
				buckets[0], buckets[1] = buckets[1], buckets[0]
			}
		} else if len(buckets) < 2 {
			buckets = append(buckets, [3]int{fruit[0], fruit[1], fruit[1]})
		} else {
			buckets = [][3]int{{buckets[1][0], buckets[1][2], buckets[1][2]}, {fruit[0], fruit[1], fruit[1]}}
		}
		if r := getFruitsNum(); result < r {
			result = r
		}
	}
	return result
}
