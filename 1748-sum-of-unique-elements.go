package main

func sumOfUnique(nums []int) int {
	// 1 <= nums[i] <= 100
	values := make([]int, 101)
	for _, num := range nums {
		values[num] += num
	}
	var result int
	for i := 1; i <= 100; i++ {
		if values[i] == i {
			result += i
		}
	}
	return result
}
