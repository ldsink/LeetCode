package main

func productExceptSelf(nums []int) []int {
	ret := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		ret[i] = 1
	}
	for i, j := 0, 1; i < len(nums)-1; i++ {
		j *= nums[i]
		ret[i+1] = j
	}
	for i, j := len(nums)-1, 1; i > 0; i-- {
		j *= nums[i]
		ret[i-1] *= j
	}
	return ret
}
