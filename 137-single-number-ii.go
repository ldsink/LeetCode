package main

import (
	"unsafe"
)

func singleNumber(nums []int) int {
	var result int

	bitLen := uint(unsafe.Sizeof(result) << 3)
	one := make([]int, bitLen)
	zero := make([]int, bitLen)

	var i uint
	for _, num := range nums {
		for i = 0; i < bitLen; i++ {
			if (num>>i)&1 == 1 {
				one[i] ++
			} else {
				zero[i] ++
			}
		}
	}

	for i = 0; i < bitLen; i++ {
		if zero[i]%3 == 0 {
			result = result ^ (1 << i)
		}
	}

	return result
}
