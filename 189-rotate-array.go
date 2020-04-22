package main

func gcd(a, b int) int {
	if a%b == 0 {
		return b
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func rotate(nums []int, k int) {
	if len(nums) == 0 {
		return
	}
	k = k % len(nums)
	if k == 0 {
		return
	}

	var index, num int
	count := lcm(len(nums), k) / k
	for c := 0; c < len(nums)/count; c++ {
		index, num = c, nums[c]
		for i := 0; i < count-1; i++ {
			l := (index + len(nums) - k) % len(nums)
			nums[index] = nums[l]
			index = l
		}
		nums[index] = num
	}
}
