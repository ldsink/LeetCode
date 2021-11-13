package main

func prefixesDivBy5(nums []int) []bool {
	var num int
	var result []bool
	for i := 0; i < len(nums); i++ {
		num <<= 1
		if nums[i] == 1 {
			num ^= 1
		}
		num %= 5
		result = append(result, num == 0)
	}
	return result
}
