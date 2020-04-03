package main

func longestConsecutive(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	valid := make(map[int]bool)
	for _, num := range nums {
		valid[num] = true
	}
	var result, left, right int
	for _, num := range nums {
		if !valid[num] {
			continue
		}
		for left = num - 1; valid[left]; left-- {
			delete(valid, left)
		}
		for right = num + 1; valid[right]; right++ {
			delete(valid, right)
		}
		if result < right-left-1 {
			result = right - left - 1
		}
	}
	return result
}
