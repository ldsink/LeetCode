package main

func _enumSubSets(nums []int, used []bool, index int) [][]int {
	if index == len(nums) {
		var result []int
		for i := 0; i < len(nums); i++ {
			if used[i] {
				result = append(result, nums[i])
			}
		}
		if len(result) > 0 {
			return [][]int{result}
		} else {
			return [][]int{}
		}
	}

	var result [][]int
	used[index] = true
	r := _enumSubSets(nums, used, index+1)
	result = append(result, r...)
	used[index] = false
	r = _enumSubSets(nums, used, index+1)
	result = append(result, r...)
	return result
}

func subsets(nums []int) [][]int {
	result := [][]int{[]int{}}
	used := make([]bool, len(nums))
	r := _enumSubSets(nums, used, 0)
	result = append(result, r...)
	return result
}
