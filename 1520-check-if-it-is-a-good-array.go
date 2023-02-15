package main

func gcd(a, b int) int {
	if a%b == 0 {
		return b
	}
	return gcd(b, a%b)
}

func isGoodArray(nums []int) bool {
	if len(nums) == 1 {
		return nums[0] == 1
	}
	x := gcd(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		x = gcd(x, nums[i])
		if x == 1 {
			return true
		}
	}
	return x == 1
}
