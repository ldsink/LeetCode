package main

func wiggleMaxLength(nums []int) int {
	if len(nums) == 1 {
		return 1
	}

	top := make([]int, len(nums))    // 当前点大于之前的点
	bottom := make([]int, len(nums)) // 当前点小于之前点
	for i := 0; i < len(nums); i++ {
		top[i] = 1
		bottom[i] = 1
		for j := i - 1; j >= 0; j-- {
			if nums[i] > nums[j] && bottom[j]+1 > top[i] {
				top[i] = bottom[j] + 1
			} else if nums[i] < nums[j] && top[j]+1 > bottom[i] {
				bottom[i] = top[j] + 1
			}
		}
	}
	var result int
	for i := len(nums) - 1; i >= 0; i-- {
		if result < top[i] {
			result = top[i]
		}
		if result < bottom[i] {
			result = bottom[i]
		}
	}
	return result
}
