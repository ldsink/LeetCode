package main

func longestEqualSubarray(nums []int, k int) int {
	type Range struct {
		start, end int
	}
	numMap := make(map[int][]Range)
	num, start := nums[0], 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != num {
			numMap[num] = append(numMap[num], Range{start: start, end: i - 1})
			num, start = nums[i], i
		}
	}
	numMap[num] = append(numMap[num], Range{start: start, end: len(nums) - 1})
	result := 0
	for _, ranges := range numMap {
		for left, right, quota, count := 0, 0, k, ranges[0].end-ranges[0].start+1; left < len(ranges); left++ {
			if left != 0 { // 不是第一个，需要退掉前面的
				quota += ranges[left].start - ranges[left-1].end - 1
				count -= ranges[left-1].end - ranges[left-1].start + 1
			}
			if right < left { // 重新初始化
				right = left
				count = ranges[left].end - ranges[left].start + 1
				quota = k
			}
			for right+1 < len(ranges) {
				if cost := ranges[right+1].start - ranges[right].end - 1; quota >= cost {
					quota -= cost
					count += ranges[right+1].end - ranges[right+1].start + 1
					right++
				} else {
					break
				}
			}
			if result < count {
				result = count
			}
		}
	}
	return result
}
