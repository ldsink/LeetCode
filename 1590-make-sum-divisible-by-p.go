package main

func minSubarray(nums []int, p int) int {
	var over int
	for _, num := range nums {
		over += num
	}
	over %= p
	if over == 0 {
		return 0
	}
	result := len(nums)
	indexes := make(map[int]int)
	indexes[0] = -1
	for i, sum := 0, 0; i < len(nums); i++ {
		sum = (sum + nums[i]) % p
		indexes[sum] = i
		target := sum - over
		if target < 0 {
			target += p
		}
		if idx, ok := indexes[target]; ok {
			if r := i - idx; result > r {
				result = r
			}
		}
	}
	if result < len(nums) {
		return result
	}
	return -1
}
