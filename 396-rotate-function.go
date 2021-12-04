package main

func maxRotateFunction(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var sum, fi int
	for i, n := range nums {
		sum += n
		fi += i * n
	}
	result := fi
	for i := 1; i < len(nums); i++ {
		fi = fi + sum - len(nums)*nums[len(nums)-i]
		if result < fi {
			result = fi
		}
	}
	return result
}
