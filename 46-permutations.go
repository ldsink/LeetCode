package main

func permute(nums []int) [][]int {
	var result [][]int
	if len(nums) == 0 {
		return result
	}
	if len(nums) == 1 {
		return append(result, []int{nums[0]})
	}

	for i := 0; i < len(nums); i++ {
		num := nums[i]
		newNums := append([]int{}, nums[:i]...)
		newNums = append(newNums, nums[i+1:]...)
		for _, r := range permute(newNums) {
			row := append([]int{num}, r...)
			result = append(result, row)
		}
	}
	return result
}
