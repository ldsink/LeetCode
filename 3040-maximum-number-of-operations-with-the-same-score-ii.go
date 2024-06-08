package main

func maxOperations(nums []int) int {
	var score, result int
	var max map[int]int
	var find func(int, int) int
	find = func(left int, right int) int {
		if left >= right {
			return 0
		}
		key := (left << 32) ^ right
		if v, ok := max[key]; ok {
			return v
		}
		if left+1 == right {
			if nums[left]+nums[right] == score {
				max[key] = 1
				return 1
			} else {
				max[key] = 0
				return 0
			}
		}
		var r int
		if nums[left]+nums[left+1] == score {
			if t := 1 + find(left+2, right); r < t {
				r = t
			}
		}
		if nums[left]+nums[right] == score {
			if t := 1 + find(left+1, right-1); r < t {
				r = t
			}
		}
		if nums[right-1]+nums[right] == score {
			if t := 1 + find(left, right-2); r < t {
				r = t
			}
		}
		max[key] = r
		return r
	}
	// 取前两个作为基准分
	score = nums[0] + nums[1]
	max = make(map[int]int)
	result = 1 + find(2, len(nums)-1)
	// 一前一后作为基准分
	score = nums[0] + nums[len(nums)-1]
	max = make(map[int]int)
	if r := 1 + find(1, len(nums)-2); result < r {
		result = r
	}
	// 取后两个作为基准分
	score = nums[len(nums)-2] + nums[len(nums)-1]
	max = make(map[int]int)
	if r := 1 + find(0, len(nums)-3); result < r {
		result = r
	}
	return result
}
