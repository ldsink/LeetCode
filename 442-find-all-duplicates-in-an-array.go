package main

// 原地替换的做法
func findDuplicates_(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			// 尝试将当前值放到正确的位置，如果无法放置，那么这个位置处理完成
			target := nums[i] - 1
			if nums[target] != target+1 {
				nums[i], nums[target] = nums[target], nums[i]
				i-- // 等下继续处理当前位置
			}
		}
	}
	var result []int
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			result = append(result, nums[i])
		}
	}
	return result
}

// 简单 bitmap 做法
func findDuplicates(nums []int) []int {
	bitmap := make([]uint64, (len(nums)/64)+1)
	var result []int
	for _, num := range nums {
		idx, offset := num/64, num%64
		if (bitmap[idx] & (1 << offset)) != 0 {
			result = append(result, num)
		} else {
			bitmap[idx] ^= 1 << offset
		}
	}
	return result
}
