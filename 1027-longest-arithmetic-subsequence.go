package main

func longestArithSeqLength(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	indexes := make(map[int][]int)
	for i, n := range nums {
		indexes[n] = append(indexes[n], i)
	}
	var result int
	for _, arr := range indexes { // 全部数值相等的等差数列
		if result < len(arr) {
			result = len(arr)
		}
	}
	getIndex := func(idx, val int) int {
		if arr, ok := indexes[val]; !ok || arr[len(arr)-1] < idx {
			return -1
		} else {
			start, end := 0, len(arr)-1
			for start != end {
				middle := (start + end) >> 1
				if idx <= arr[middle] {
					end = middle
				} else {
					start = middle + 1
				}
			}
			return arr[start]
		}
	}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			l, delta := 2, nums[j]-nums[i]
			if delta == 0 {
				continue
			}
			for val, idx := nums[j]+delta, j+1; ; l++ {
				if idx = getIndex(idx, val); idx != -1 {
					val, idx = val+delta, idx+1
				} else {
					break
				}
			}
			if result < l {
				result = l
			}
		}
	}
	return result
}
