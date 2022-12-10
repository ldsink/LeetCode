package main

func maxJump(stones []int) int {
	if len(stones) == 2 {
		return stones[1]
	}
	var result int
	for i := 0; i < len(stones)-2; i++ {
		if l := stones[i+2] - stones[i]; result < l {
			result = l
		}
	}
	return result
}
