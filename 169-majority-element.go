package main

func majorityElement(nums []int) int {
	var major, count int
	for i := 0; i+1 < len(nums); i += 2 {
		if nums[i] != nums[i+1] {
			// pass
		} else if count == 0 {
			count++
			major = nums[i]
		} else if major != nums[i] {
			count--
		} else {
			count++
		}
	}
	if count > 0 {
		return major
	} else {
		return nums[len(nums)-1]
	}
}
