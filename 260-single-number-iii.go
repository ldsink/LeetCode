package main

import (
	"math/bits"
)

func singleNumber(nums []int) []int {
	if len(nums) == 2 {
		return nums
	}

	xor := nums[0]
	for i := 1; i < len(nums); i++ {
		xor ^= nums[i]
	}
	offset := uint(bits.Len(uint(xor)) - 1)
	var zero, one int
	for _, num := range nums {
		if (num>>offset)&0x1 == 0 {
			zero ^= num
		} else {
			one ^= num
		}
	}
	return []int{zero, one}
}
