package main

import (
	"math/bits"
)

func bucketSort(nums []int, length uint) []int {
	var zero, one []int
	for _, num := range nums {
		if (num>>length)&0x1 == 1 {
			one = append(one, num)
		} else {
			zero = append(zero, num)
		}
	}
	if length == 0 {
		return append(zero, one...)
	}
	if len(zero) == 0 || len(one) == 0 {
		return bucketSort(nums, length-1)
	}
	return append(bucketSort(zero, length-1), bucketSort(one, length-1)...)
}

func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}

	// actually, sort.Ints is faster than bucketSort
	//sort.Ints(nums)
	length := 1
	for _, num := range nums {
		if l := bits.Len(uint(num)); length < l {
			length = l
		}
	}
	nums = bucketSort(nums, uint(length-1))

	var result int
	for i := 0; i < len(nums)-1; i++ {
		if r := nums[i+1] - nums[i]; result < r {
			result = r
		}
	}
	return result
}
