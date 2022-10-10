package main

import "math"

func minSwap(nums1 []int, nums2 []int) int {
	nums := []*[]int{&nums1, &nums2}
	count := [2][2]int{{0, 0}, {1, 0}}
	for i := 1; i < len(nums1); i++ {
		for j := 0; j < 2; j++ {
			swap := math.MaxInt
			for k := 0; k < 2; k++ {
				if (*nums[j])[i] > (*nums[k])[i-1] && (*nums[j^1])[i] > (*nums[k^1])[i-1] && swap > count[k][(i-1)%2]+j {
					swap = count[k][(i-1)%2] + j
				}
			}
			if swap == math.MaxInt {
				swap--
			}
			count[j][i%2] = swap
		}
	}
	idx := (len(nums1) - 1) % 2
	result := count[0][idx]
	if result > count[1][idx] {
		result = count[1][idx]
	}
	return result
}
