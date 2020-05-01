package main

func find(k, v, n int, used *[10]bool, result *[][]int) {
	if k == 0 || n == 0 {
		if k == 0 && n == 0 {
			var nums []int
			for i := 1; i < 10; i++ {
				if (*used)[i] {
					nums = append(nums, i)
				}
			}
			*result = append(*result, nums)
		}
		return
	}

	if v+k == 10 {
		if n == 45-v*(v-1)/2 {
			var nums []int
			for i := 1; i < v; i++ {
				if (*used)[i] {
					nums = append(nums, i)
				}
			}
			for i := v; i < 10; i++ {
				nums = append(nums, i)
			}
			*result = append(*result, nums)
		}
		return
	}

	(*used)[v] = true
	find(k-1, v+1, n-v, used, result)
	(*used)[v] = false
	find(k, v+1, n, used, result)
}

func combinationSum3(k int, n int) [][]int {
	var result [][]int
	if k > 9 || n > 45 {
		return result
	}
	var used [10]bool
	find(k, 1, n, &used, &result)
	return result
}
