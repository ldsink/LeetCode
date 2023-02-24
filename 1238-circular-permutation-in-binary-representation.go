package main

func circularPermutation(n int, start int) []int {
	nums := []int{0, 1}
	for i := 1; i < n; i++ {
		newNums := append([]int{}, nums...)
		for j := len(nums) - 1; j >= 0; j-- {
			newNums = append(newNums, 1<<i+nums[j])
		}
		nums = newNums
	}
	var idx int
	for i, num := range nums {
		if num == start {
			idx = i
			break
		}
	}
	var result []int
	for i := 0; i < len(nums); i++ {
		result = append(result, nums[(i+idx)%len(nums)])
	}
	return result
}
