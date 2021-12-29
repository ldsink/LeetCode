package main

func countQuadruplets(nums []int) int {
	var result int
	for i := 0; i < len(nums)-3; i++ {
		for j := i + 1; j < len(nums)-2; j++ {
			for k := j + 1; k < len(nums)-1; k++ {
				for l := k + 1; l < len(nums); l++ {
					if nums[l] == nums[i]+nums[j]+nums[k] {
						result++
					}
				}
			}
		}
	}
	return result
}
