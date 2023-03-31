package main

func arithmeticTriplets(nums []int, diff int) int {
	exists := make([]bool, 201)
	for _, num := range nums {
		exists[num] = true
	}
	var result int
	for i := 0; i < len(nums)-2; i++ {
		if nums[i]+diff<<1 >= len(exists) {
			break
		}
		if exists[nums[i]+diff] && exists[nums[i]+diff<<1] {
			result++
		}
	}
	return result
}
