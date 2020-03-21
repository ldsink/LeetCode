package main

func jump(nums []int) int {
	minJumps := make([]int, len(nums))
	if len(minJumps) > 0 {
		minJumps[len(nums)-1] = 0
	}
	for i := len(nums) - 2; i >= 0; i-- {
		j := i + nums[i]
		if j >= len(nums)-1 {
			minJumps[i] = 1
			continue
		}
		minJumps[i] = minJumps[i+1] + 1
		for k := i + 2; k <= j; k++ {
			if minJumps[i] > minJumps[k]+1 {
				minJumps[i] = minJumps[k] + 1
			}
		}
	}
	return minJumps[0]
}
