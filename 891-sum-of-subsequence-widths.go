package main

import "sort"

func sumSubseqWidths(nums []int) int {
	const mod = 1e9 + 7
	sort.Ints(nums)
	result, a, b := 0, nums[0], 2
	for i := 1; i < len(nums); i++ {
		result += nums[i]*(b-1) - a
		result %= mod
		a = (a<<1 + nums[i]) % mod
		b = b << 1 % mod
	}
	for result < 0 {
		result += mod
	}
	return result % mod
}
