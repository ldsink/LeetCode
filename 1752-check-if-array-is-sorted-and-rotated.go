package main

func check(nums []int) bool {
	idx := -1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			idx = i + 1
			break
		}
	}
	if idx == -1 {
		return true
	}
	for i := 0; i < len(nums)-1; i++ {
		j := (idx + i) % len(nums)
		k := (j + 1) % len(nums)
		if nums[j] > nums[k] {
			return false
		}
	}
	return true
}
