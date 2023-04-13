package main

func mostFrequentEven(nums []int) int {
	count := make(map[int]int)
	var maxCount, minNum int
	for _, num := range nums {
		if num&1 == 1 {
			continue
		}
		count[num]++
		if maxCount < count[num] || (maxCount == count[num] && num < minNum) {
			maxCount, minNum = count[num], num
		}
	}
	if maxCount == 0 {
		return -1
	}
	return minNum
}
