package main

func numSubarrayBoundedMax(nums []int, left int, right int) int {
	valid, candidate := make([]int32, len(nums)), make([]int32, len(nums))
	for idx, num := range nums {
		if num > right {
			continue
		} else if num < left {
			if idx > 0 {
				valid[idx], candidate[idx] = valid[idx-1], candidate[idx-1]+1
			} else {
				valid[idx], candidate[idx] = 0, 1
			}
		} else {
			if idx > 0 {
				valid[idx], candidate[idx] = 1+valid[idx-1]+candidate[idx-1], 0
			} else {
				valid[idx], candidate[idx] = 1, 0
			}
		}
	}
	var result int32
	for i := 0; i < len(nums); i++ {
		result += valid[i]
	}
	return int(result)
}
