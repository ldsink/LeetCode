package main

func numSubarrayProductLessThanK(nums []int, k int) int {
	var start, end, multi, result int
	for start, multi = 0, 1; start < len(nums); start++ {
		if end < start { // 当不存在连续的情况，初始化 end 和 multi 的值
			end, multi = start, 1
		}
		for ; end < len(nums); end++ {
			if multi*nums[end] < k {
				multi *= nums[end] // 使用 multi 来累积 [start, end) 这一段的乘积，避免重复计算
			} else {
				break
			}
		}
		result += end - start     // 以 start 为起点的全部满足条件的子数组的数量
		if multi >= nums[start] { // 从 multi 移除起点的乘积
			multi /= nums[start]
		}
	}
	return result
}
